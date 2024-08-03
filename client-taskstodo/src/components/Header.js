import logo from "../img/logo.svg";
import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import { Link } from "react-router-dom";
import { useState } from "react";
import CreateTaskModal from "../components/CreateTaskModal";

const Header = (props) => {
    const [showModal, setShowModal] = useState(false);

    const logOut = () => {
        props.setJwtToken("");
    };

    const handleTaskCreated = () => {
        setShowModal(false);
        // Notificar o App sobre a criação de uma nova tarefa.
        props.onTaskCreated();
    };

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
                            {props.jwtToken ? (
                                <>
                                    <Nav.Link as={Link} to="#" onClick={() => setShowModal(true)}>+ Create Task</Nav.Link>
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

            <CreateTaskModal
                show={showModal}
                onClose={() => setShowModal(false)}
                onTaskCreated={handleTaskCreated}
            />
        </>
    );
};

export default Header;
