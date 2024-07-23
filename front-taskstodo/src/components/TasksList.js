import { useEffect, useState } from "react";
import { Link, useOutletContext } from "react-router-dom";


const TaskList = () => {
	const { user } = useOutletContext(); // get user mock login context 
	const [tasks, setTasks] = useState([]);


	useEffect(() => {
		if (user) { // if the user in logged

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


	}, [user]);


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
			<div className="col">
				{!user ?
					(<div className="alert alert-warning" role="alert">
						Please log in to see your tasks.
					</div>)

					: (tasks.map((listItem) => (
						<Link key={listItem.id} to={`/tasks/${listItem.id}`} className="text-decoration-none">
							<div className="card mb-3" style={{ transition: 'transform 0.2s', boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)' }}
								onMouseOver={e => e.currentTarget.style.transform = 'translateY(-5px)'}
								onMouseOut={e => e.currentTarget.style.transform = 'translateY(0)'}>
								<div className="card-body">
									<h5 className="card-title">{listItem.title}</h5>
									<p className="card-text">{listItem.description}</p>
									<span className={`badge ${getStatusBadgeClass(listItem.status)}`}>
										Status: {listItem.status}
									</span>
								</div>
							</div>

						</Link>
					))

					)
				}

			</div>
		</>
	);
};

export default TaskList;

