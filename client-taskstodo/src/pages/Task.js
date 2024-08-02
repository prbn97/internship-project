import { useEffect, useState } from "react";
import { useParams, useOutletContext, useNavigate } from "react-router-dom";


// in this components we display the task by id
const Task = () => {
    const [task, setTask] = useState({}); // 1ª set state with empty object to store the task
    // 2ª get the id from the url to do the request
    // using the hook useParams to make this happen
    let { id } = useParams(); // this variable needs to match with the :id in index.js

    const { jwtToken } = useOutletContext();
    const navigate = useNavigate();
    // 3ª useEffect to make the request
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
                })
                .catch(err => {
                    console.log(err);
                });

        }
    }, [id, jwtToken, navigate]); // pass the id

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

    return (
        <div className="card mb-3" style={{ transition: 'transform 0.2s', boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)' }}
            onMouseOver={e => e.currentTarget.style.transform = 'translateY(-5px)'}
            onMouseOut={e => e.currentTarget.style.transform = 'translateY(0)'}>
            <div className="card-body">
                <h5 className="card-title">{task.title}</h5>
                <p className="card-text">{task.description}</p>
                <div className="mb-3">
                    <button
                        className={`btn ${getStatusBadgeClass(task.status)} cursor-pointer`}
                    // onClick={handleStatusClick}
                    > Status: {task.status}
                    </button>
                </div>
            </div>
        </div>
    )
}

export default Task;