package nmap

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/OVantsevich/security-utils/server/internal/model"
	"golang.org/x/net/context"
)

type cacheItem struct {
	filename string
	ttl      time.Time
}

type Nmap struct {
	cache map[string]cacheItem

	ttl      time.Duration
	filesDir string
}

func NewNmap(ttl time.Duration, filesDir string) *Nmap {
	return &Nmap{cache: make(map[string]cacheItem), ttl: ttl, filesDir: filesDir}
}

func (n *Nmap) Get(ctx context.Context, parameters model.NmapScanParameters) ([]byte, error) {
	item, ok := n.cache[hash(parameters)]
	if !ok {
		_, err := n.update(ctx, parameters)
		if err != nil {
			return nil, fmt.Errorf("n.update: %w", err)
		}
		item = cacheItem{
			filename: hash(parameters),
			ttl:      time.Now().UTC().Add(n.ttl),
		}
		n.cache[hash(parameters)] = item
	}
	xmlFilename := fmt.Sprintf("%s.xml", n.filename(item.filename))
	output, err := xsltproc(ctx, xmlFilename)
	if err != nil {
		return nil, fmt.Errorf("cmd.CombinedOutput: %w", err)
	}

	return output, nil
}

func (n *Nmap) GetXML(ctx context.Context, parameters model.NmapScanParameters) ([]byte, error) {
	if item, ok := n.cache[hash(parameters)]; ok && time.Now().Before(item.ttl) {
		xmlFilename := fmt.Sprintf("%s.xml", item.filename)

		data, err := os.ReadFile(n.filesDir + xmlFilename)
		if err != nil {
			return nil, fmt.Errorf("os.ReadFile: %w", err)
		}

		return data, nil
	}

	data, err := n.update(ctx, parameters)
	if err != nil {
		return nil, fmt.Errorf("n.update: %w", err)
	}
	return data, nil
}

func (n *Nmap) update(ctx context.Context, parameters model.NmapScanParameters) ([]byte, error) {
	data, err := nmap(ctx, parameters)
	if err != nil {
		return nil, err
	}

	filename := hash(parameters)
	xmlFilename := fmt.Sprintf("%s.xml", filename)

	err = n.setXMLFile(xmlFilename, data)
	if err != nil {
		return nil, fmt.Errorf("n.setXMLFile: %w", err)
	}

	n.cache[hash(parameters)] = cacheItem{
		filename: filename,
		ttl:      time.Now().UTC().Add(n.ttl),
	}

	return data, nil
}

func (n *Nmap) setXMLFile(filename string, data []byte) error {
	file, err := os.Create(n.filename(filename))
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("file.Write: %w", err)
	}
	return nil
}

func (n *Nmap) filename(filename string) string {
	return fmt.Sprintf("%s/%s", n.filesDir, filename)
}

func nmap(ctx context.Context, parameters model.NmapScanParameters) ([]byte, error) {
	cmdArgs := parameters.ToCmd()
	cmd := exec.CommandContext(ctx, "nmap", cmdArgs...)

	buf := &bytes.Buffer{}
	cmd.Stdout = buf
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("cmd.Run: %w", err)
	}
	return buf.Bytes(), nil
}

func xsltproc(ctx context.Context, fileName string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "xsltproc", fileName)

	buf := &bytes.Buffer{}
	cmd.Stdout = buf
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("cmd.Run: %w", err)
	}
	return buf.Bytes(), nil
}

func hash(s model.NmapScanParameters) string {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(s)
	d := sha256.Sum256(b.Bytes())
	return base64.URLEncoding.EncodeToString(d[:])
}
