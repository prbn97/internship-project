import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { Container, Nav, Navbar, Modal, Button, Form } from 'react-bootstrap';
import logo from './img/logo.svg';
import create from './img/create.svg';
import login from './img/login.svg';
import Title from './Title';

function Header({ jwtToken, setJwtToken, user, setUser }) {
    const [showModal, setShowModal] = useState(false);
    const [title, setTitle] = useState("");
    const [description, setDescription] = useState("");
    const [errors, setErrors] = useState({});
    const navigate = useNavigate();

    const handleCloseModal = () => {
        setShowModal(false);
        setTitle("");
        setDescription("");
        setErrors({});
    };
    const handleOpenModal = () => setShowModal(true);

    const handleLogout = () => {
        setJwtToken("");
        setUser(null);
    };

    const validateForm = () => {
        const newErrors = {};
        if (!title.trim()) {
            newErrors.title = "Title cannot be empty or just spaces.";
        }
        return newErrors;
    };

    const handlePOST = (title, description) => {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const requestOptions = {
            method: "POST",
            headers: headers,
            body: JSON.stringify({
                title: title,
                description: description,
            }),
        };

        fetch(`http://localhost:8080/tasks`, requestOptions)
            .then((response) => {
                if (response.ok) {
                    navigate('/tasks');
                }
            })
            .catch(err => {
                console.log(err);
            })
    };

    const handleSubmit = () => {
        const validationErrors = validateForm();
        if (Object.keys(validationErrors).length > 0) {
            setErrors(validationErrors);
        } else {
            handlePOST(title, description)
            handleCloseModal();
        }
    };

    return (
        <>
            <Navbar expand="lg" className="bg-body-tertiary">
                <Container className="d-flex justify-content-between align-items-center">
                    <Navbar.Brand as={Link} to="/" className="d-flex align-items-center">
                        <img
                            alt=""
                            src={logo}
                            width="30"
                            height="30"
                            className="d-inline-block align-top"
                        />
                        <h4 className="mb-0 me-2">
                            Tasks to<strong className='text-Dark'>Do</strong>
                        </h4>
                    </Navbar.Brand>

                    <Navbar.Toggle aria-controls="basic-navbar-nav" />
                    <Navbar.Collapse id="basic-navbar-nav" className="d-flex justify-content-end">
                        <Nav className="d-flex align-items-center">
                            <Nav.Link
                                onClick={handleOpenModal}
                                className="d-flex align-items-center"
                                disabled={jwtToken === ""}
                            >
                                <img alt="" src={create} width="30" height="30"
                                    className="d-inline-block align-top" />
                                Create Task
                            </Nav.Link>
                            {jwtToken === ""
                                ? <Nav.Link as={Link} to="/login" className="d-flex align-items-center">
                                    Login
                                    <img alt="" src={login} width="30" height="30"
                                        className="d-inline-block align-top m-2"
                                    />
                                </Nav.Link>
                                : <Nav.Link as={Link} to="/" onClick={handleLogout} className="d-flex align-items-center">
                                    Logout
                                    <img alt="" src={login} width="30" height="30"
                                        className="d-inline-block align-top m-2"
                                    />
                                </Nav.Link>
                            }
                        </Nav>
                    </Navbar.Collapse>
                </Container>
            </Navbar >

            <Modal show={showModal} onHide={handleCloseModal}>
                <Modal.Header closeButton>
                    <Modal.Title>
                        <Title icon={logo} text="Create Task" />
                    </Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <Form>
                        <Form.Group controlId="formTaskTitle">
                            <Form.Label>Title</Form.Label>
                            <Form.Control
                                type="text"
                                value={title}
                                onChange={(e) => setTitle(e.target.value)}
                                isInvalid={!!errors.title}
                            />
                            <Form.Control.Feedback type="invalid">
                                {errors.title}
                            </Form.Control.Feedback>
                        </Form.Group>
                        <Form.Group controlId="formTaskDescription" className="mt-3">
                            <Form.Label>Description</Form.Label>
                            <Form.Control
                                as="textarea"
                                rows={3}
                                value={description}
                                onChange={(e) => setDescription(e.target.value)}
                            />
                        </Form.Group>

                    </Form>
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary" onClick={handleCloseModal}>
                        Close
                    </Button>
                    <Button variant="primary" onClick={handleSubmit}>
                        Create
                    </Button>
                </Modal.Footer>
            </Modal>
        </>
    );
}

export default Header;
