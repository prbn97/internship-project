import logo from "../img/logo.svg"
import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';

import { Link } from "react-router-dom";

const Header = (props) => {

    const logOut = () => {
        props.setJwtToken("")
    }


    return (
        <>
            <Navbar className="bg-body-tertiary">
                <Container >
                    <Navbar.Brand as={Link} to="/">
                        <img
                            alt="logo"
                            src={logo}
                            width="30"
                            height="30"
                            className="d-inline-block align-top"
                        />{' '}
                        Tasks ToDo
                    </Navbar.Brand>

                    <Navbar.Collapse id="basic-navbar-nav" className="flex-row-reverse">
                        <Nav>
                            {/* using "as={Link} to=" to use React Router's Link*/}
                            <Nav.Link as={Link} to="/">Home</Nav.Link>
                            {props.jwtToken ? (
                                <>
                                    <Nav.Link as={Link} to="/tasks">Tasks</Nav.Link>
                                    <Nav.Link as={Link} to="/login" onClick={logOut}>Logout</Nav.Link>
                                </>
                            ) : (
                                <Nav.Link as={Link} to="/login">Login</Nav.Link>
                            )}
                        </Nav>
                    </Navbar.Collapse>
                </Container>
            </Navbar>

        </>
    );
}

export default Header;