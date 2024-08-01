import { useEffect, useState } from "react";

import TaskCard from "../components/TaskCard";


// in this components we display the user's tasks
// for this we call the API to get all tasks
const TasksList = () => {
    //  using React State with empty array for store tasks
    const [tasks, setTasks] = useState([])

    // when this component load useEffect hook make the request to the API
    useEffect(() => {
        let tasksList = [
            // create this fake tasks for now
            {
                "id": "1",
                "title": "task 1",
                "description": "",
                "status": "ToDo"
            },
            {
                "id": "2",
                "title": "task 2",
                "description": "",
                "status": "ToDo"
            },
        ];
        // then set the http response to setTasks
        setTasks(tasksList)

    }, /* to just run once we do it like that --> */[]);
    return (
        <>
            <div className="col mt-2">
                {tasks.map((task) => (<TaskCard key={task.id} task={task} />)
                )}
            </div>
        </>
    )
}

export default TasksList;