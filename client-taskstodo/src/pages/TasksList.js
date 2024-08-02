import TaskCard from "../components/TaskCard";

import { useEffect, useState } from "react";
import { useNavigate, useOutletContext } from "react-router-dom";


// in this components we display the user's tasks
// for this we call the API to get all tasks
const TasksList = () => {
    //  using React State with empty array for store tasks
    const [tasks, setTasks] = useState([])
    const { jwtToken, setAlertClassName, setAlertMessage } = useOutletContext()
    const navigate = useNavigate();

    // when this component load useEffect hook make the request to the API
    useEffect(() => {
        if (!jwtToken) {
            navigate("/login");
            return;
        } else {
            const headers = new Headers();
            headers.append("Content-Type", "application/json");

            // get tasks list
            const requestOptions = {
                method: "GET",
                headers: headers,
            };

            fetch(`${process.env.REACT_APP_BACKEND_ADDRESS}/tasks`, requestOptions)
                .then((response) => response.json())
                .then((data) => {
                    setTasks(data);
                    setAlertClassName("d-none")
                    setAlertMessage("")
                })
                .catch(err => {
                    console.log(err);
                    setAlertClassName("alert-danger")
                    setAlertMessage("internal server error")
                });

        };
    }, [jwtToken, navigate, setAlertClassName, setAlertMessage]);

    return (
        <>
            <div className="col mt-2">
                {tasks.map((task) => (
                    <TaskCard key={task.id} task={task} />
                ))}
            </div>
        </>
    );
}

export default TasksList;