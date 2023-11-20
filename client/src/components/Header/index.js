import React, {useState} from 'react';
import styled from 'styled-components';
import { Link } from 'react-router-dom';

import logo from "../../assets/logo.png";

const Container = styled.nav`
  background-color: DarkBlue;
`;

const Logo = styled.img`
  height: 60px;
  width: 60px;
  margin-right: 16px;
`;

const NavItem = styled.li`
  a {
    color: white !important;
    font-size: 20px;
    padding: 10px;
     border-radius: 5px;
    transition: background-color 0.3s;

    &.active {
      background-color: #007bff;
      color: white;
    }
  }
`;

function Header() {
    const [activeItem, setActiveItem] = useState('default'); // Устанавливаем тему "default" по умолчанию

    const handleItemClick = (item) => {
        setActiveItem(item);
    };

    return (
        <Container className="navbar navbar-expand-lg">
            <div className="container-fluid">
                <a className="navbar-brand" href="#">
                    <Logo src={logo} alt="Logo"/>
                </a>
                <button className="navbar-toggler" type="button" data-bs-toggle="collapse"
                        data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent"
                        aria-expanded="false"
                        aria-label="Toggle navigation">
                    <span className="navbar-toggler-icon"></span>
                </button>
                <div className="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul className="navbar-nav me-auto mb-2 mb-lg-0">
                        <NavItem className="nav-item">
                            <Link
                                to="/gobuster" // Используем Link для перехода на страницу "gobuster"
                                className={`nav-link ${activeItem === 'gobuster' ? 'active' : ''}`}
                            >
                                gobuster
                            </Link>
                        </NavItem>
                        <NavItem className="nav-item">
                            <Link
                                to="/nmap" // Используем Link для перехода на страницу "nmap"
                                className={`nav-link ${activeItem === 'nmap' ? 'active' : ''}`}
                            >
                                nmap
                            </Link>
                        </NavItem>
                    </ul>
                </div>
            </div>
        </Container>
    );
}

export default Header;
