import { useEffect } from "react";
import { Link, useOutletContext } from "react-router-dom";
import TaskItem from "./TaskItem";

const TaskList = () => {
	const { user, tasks, setTasks } = useOutletContext(); // get user context

	useEffect(() => {
		if (user) { // if the user is logged in
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
				})
				.catch(err => {
					console.log(err);
				});
		}
	}, [user, setTasks]);

	const getStatusBadgeClass = (status) => {
		switch (status) {
			case "ToDo":
				return "bg-danger";
			case "Doing":
				return "bg-warning";
			case "Done":
				return "bg-success";
			default:
				return "bg-secondary";
		}
	};

	return (
		<>
			<div className="col mt-2">
				{!user ?
					(<div className="alert alert-warning" role="alert">
						Please log in to see your tasks.
					</div>)
					: (tasks.map((task) => (<>
						<TaskItem task={task} />

					</>
					)))
				}
			</div>
		</>
	);
};

export default TaskList;
