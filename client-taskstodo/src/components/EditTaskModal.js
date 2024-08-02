import React, { useState } from 'react';
import { Modal, Button, Form } from 'react-bootstrap';

const EditTaskModal = ({ show, handleClose, handleSave, task }) => {
    const [title, setTitle] = useState(task.title);
    const [description, setDescription] = useState(task.description);

    const handleSubmit = (e) => {
        e.preventDefault();
        handleSave({ title, description });
    };

    return (
        <Modal show={show} onHide={handleClose}>
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
                    <Button variant="primary" type="submit">
                        Save Changes
                    </Button>
                </Form>
            </Modal.Body>
        </Modal>
    );
};

export default EditTaskModal;
