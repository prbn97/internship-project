import { useEffect, useState } from "react";
import { useNavigate, useOutletContext } from "react-router-dom";
import TaskCard from "../components/TaskCard";

const TasksList = () => {
    const [tasks, setTasks] = useState([]);
    const { jwtToken, setAlertClassName, setAlertMessage, tasksUpdate } = useOutletContext();
    const navigate = useNavigate();

    useEffect(() => {
        if (!jwtToken) {
            navigate("/login");
            return;
        }

        const fetchTasks = () => {
            const headers = new Headers();
            headers.append("Content-Type", "application/json");

            const requestOptions = {
                method: "GET",
                headers: headers,
            };

            fetch(`${process.env.REACT_APP_BACKEND_ADDRESS}/tasks`, requestOptions)
                .then((response) => response.json())
                .then((data) => {
                    setTasks(data);
                    setAlertClassName("d-none");
                    setAlertMessage("");
                })
                .catch(err => {
                    console.log(err);
                    setAlertClassName("alert-danger");
                    setAlertMessage("internal server error");
                });
        };

        fetchTasks();
    }, [jwtToken, navigate, setAlertClassName, setAlertMessage, tasksUpdate]);

    return (
        <div className="col mt-2">
            {tasks.map((task) => (
                <TaskCard key={task.id} task={task} />
            ))}
        </div>
    );
};

export default TasksList;
