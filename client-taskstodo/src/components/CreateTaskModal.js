import { useState, useEffect } from "react";
import { Modal, Button, Form } from "react-bootstrap";

const CreateTaskModal = ({ show, onClose, onTaskCreated }) => {
    const [title, setTitle] = useState("");
    const [description, setDescription] = useState("");
    const [error, setError] = useState("");

    const handleSubmit = (event) => {
        event.preventDefault();
        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        if (!title.trim()) {
            setError("Title cannot be empty.");
            return;
        }
        setError("");

        const requestOptions = {
            method: "POST",
            headers: headers,
            body: JSON.stringify({ title, description }),
        };

        fetch(`${process.env.REACT_APP_BACKEND_ADDRESS}/tasks`, requestOptions)
            .then((response) => response.json())
            .then(() => {
                setTitle("");
                setDescription("");
                onTaskCreated(); // Chama a função para atualizar a lista de tarefas
            })
            .catch(err => {
                console.log(err);
            });
    };

    // Clear error message when modal is closed
    useEffect(() => {
        if (!show) {
            setError("");
        }
    }, [show]);

    return (
        <Modal show={show} onHide={onClose}>
            <Modal.Header closeButton>
                <Modal.Title>Create New Task</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form onSubmit={handleSubmit}>
                    <Form.Group controlId="formTaskTitle">
                        <Form.Label>Title</Form.Label>
                        <Form.Control
                            type="text"
                            placeholder="Enter task title"
                            value={title}
                            onChange={(e) => setTitle(e.target.value)}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="formTaskDescription">
                        <Form.Label>Description</Form.Label>
                        <Form.Control
                            as="textarea"
                            rows={3}
                            placeholder="Enter task description"
                            value={description}
                            onChange={(e) => setDescription(e.target.value)}
                        />
                    </Form.Group>
                    {error && <div className="alert alert-danger">{error}</div>}
                    <Button variant="primary" type="submit">
                        Create Task
                    </Button>
                </Form>
            </Modal.Body>
        </Modal>
    );
};

export default CreateTaskModal;
