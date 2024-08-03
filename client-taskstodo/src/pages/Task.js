import editIcon from '../img/edit.svg';
import deleteIcon from '../img/delete.svg';
import DeleteTaskModal from '../components/DeleteTaskModal';
import EditTaskModal from '../components/EditTaskModal';

import { useEffect, useState } from "react";
import { useParams, useOutletContext, useNavigate } from "react-router-dom";


// in this components we display the task by id
const Task = () => {
    const [task, setTask] = useState({}); // set state with empty object to store the task
    // get the id from the url to do the request
    // using the hook useParams to make this happen
    let { id } = useParams(); // this variable needs to match with the :id in index.js

    const { jwtToken, setAlertClassName, setAlertMessage } = useOutletContext()
    const navigate = useNavigate();

    useEffect(() => {
        if (!jwtToken) {
            navigate("/login");
            return;
        } else {
            const headers = new Headers();
            headers.append("Content-Type", "application/json");

            const requestOptions = {
                method: "GET",
                headers: headers,
            }; // get task by id
            fetch(`${process.env.REACT_APP_BACKEND_ADDRESS}/tasks/${id}`, requestOptions)
                .then((response) => response.json())
                .then((data) => {
                    setTask(data);
                    setAlertClassName("d-none")
                    setAlertMessage("")
                })
                .catch(err => {
                    console.log(err);
                    setAlertClassName("alert-danger")
                    setAlertMessage("internal server error")
                });

        }
    }, [id, jwtToken, navigate, setAlertClassName, setAlertMessage]); // pass the id

    const getStatusBadgeClass = (status) => {
        switch (status) {
            case "ToDo":
                return "btn-danger";
            case "Doing":
                return "btn-warning";
            case "Done":
                return "btn-success";
            default:
                return "btn-secondary";
        }
    };
    const handleStatusClick = () => {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const requestOptions = {
            method: "PUT",
            headers: headers,
        };

        fetch(`${process.env.REACT_APP_BACKEND_ADDRESS}/tasks/${id}/update`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                setTask(data);
            })
            .catch(err => {
                console.log(err);
            });
    };

    const [showDeleteModal, setShowDeleteModal] = useState(false)
    const handleDelete = () => {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const requestOptions = {
            method: "DELETE",
            headers: headers,
        };

        fetch(`${process.env.REACT_APP_BACKEND_ADDRESS}/tasks/${id}`, requestOptions)
            .then((response) => {
                if (response.ok) {
                    navigate("/tasks");
                } else {
                    throw new Error("Failed to delete task");
                }
            })
            .catch(err => {
                console.log(err);
                setAlertClassName("alert-danger");
                setAlertMessage("Failed to delete task");
            });
        setShowDeleteModal(false);
    };

    const [showEditModal, setShowEditModal] = useState(false);
    const handleSaveEdit = (updatedTask) => {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const requestOptions = {
            method: "PUT",
            headers: headers,
            body: JSON.stringify(updatedTask),
        };

        fetch(`${process.env.REACT_APP_BACKEND_ADDRESS}/tasks/${id}`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                setTask(data);
                setShowEditModal(false);
            })
            .catch(err => {
                console.log(err);
                setAlertClassName("alert-danger");
                setAlertMessage("Failed to update task");
            });
    };
    return (
        <div className="card mt-3" style={{ boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)' }}>
            <div className="card-body">
                <div className="d-flex justify-content-between">
                    <h5 className="card-title">{task.title}</h5>
                    <img
                        alt="delete-task-icon" src={editIcon}
                        width="30" height="30" className="ms-5 cursor-pointer"
                        onClick={() => setShowEditModal(true)}
                    />
                </div>
                <p className="card-text">{task.description}</p>
                <div className="d-flex justify-content-between">
                    <button
                        className={`btn ${getStatusBadgeClass(task.status)} cursor-pointer`}
                        onClick={handleStatusClick}> Status: {task.status}
                    </button>
                    <img
                        alt="delete-task-icon" src={deleteIcon}
                        width="30" height="30" className="ms-5 cursor-pointer"
                        onClick={() => setShowDeleteModal(true)}
                    />
                </div>
                <DeleteTaskModal show={showDeleteModal} handleClose={() => setShowDeleteModal(false)} handleDelete={handleDelete} />
                <EditTaskModal show={showEditModal} handleClose={() => setShowEditModal(false)} handleSave={handleSaveEdit} task={task} />
            </div>
        </div>
    )
}

export default Task;