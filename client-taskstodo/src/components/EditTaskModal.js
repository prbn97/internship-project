import React, { useState, useEffect } from 'react';
import { Modal, Button, Form } from 'react-bootstrap';

const EditTaskModal = ({ show, handleClose, handleSave, task }) => {
    const [title, setTitle] = useState("");
    const [description, setDescription] = useState("");
    const [error, setError] = useState("");

    // Define title and description when task prop changes
    useEffect(() => {
        if (task) {
            setTitle(task.title || "");
            setDescription(task.description || "");
        }
    }, [task]);

    // Clear error message when modal is closed
    useEffect(() => {
        if (!show) {
            setError("");
        }
    }, [show]);

    const handleSubmit = (e) => {
        e.preventDefault();

        if (!title.trim()) {
            setError("Title cannot be empty.");
            return;
        }
        setError(""); // Clear error message on successful submission
        handleSave({ title, description });
    };

    return (
        <Modal show={show} onHide={() => {
            handleClose();
            setError(""); // Clear error message when modal is closed
        }}>
            <Modal.Header closeButton>
                <Modal.Title>Edit Task</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form onSubmit={handleSubmit}>
                    <Form.Group className="mb-3" controlId="formTaskTitle">
                        <Form.Label>Title</Form.Label>
                        <Form.Control
                            type="text"
                            value={title}
                            onChange={(e) => setTitle(e.target.value)}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="formTaskDescription">
                        <Form.Label>Description</Form.Label>
                        <Form.Control
                            as="textarea"
                            rows={3}
                            value={description}
                            onChange={(e) => setDescription(e.target.value)}
                        />
                    </Form.Group>
                    {error && <div className="alert alert-danger">{error}</div>}
                    <Button variant="primary" type="submit">
                        Save Changes
                    </Button>
                </Form>
            </Modal.Body>
        </Modal>
    );
};

export default EditTaskModal;
