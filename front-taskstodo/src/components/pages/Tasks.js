import { useEffect, useState } from "react";
import logo from '../img/logo.svg';
import Title from '../Title';
import { Link, useOutletContext } from "react-router-dom";



const Tasks = () => {
    const [tasks, setTasks] = useState([]);
    const { user } = useOutletContext(); // get user context 

    useEffect(() => {
        if (user) { // mock login
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
                return "bg-secondary"; // fallback color
        }
    };

    return (
        <>
            <div className="container">
                <div className="row justify-content-center">
                    <div className="col-auto">
                        <Title icon={logo} text="List of Tasks" />
                    </div>
                </div>
                <div className="row m-4">
                    <div className="col">
                        {!user ? (
                            <div className="alert alert-warning" role="alert">
                                Please log in to see your tasks.
                            </div>
                        ) : (
                            tasks.map((listItem) => (
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
                        )}
                    </div>
                </div>
            </div>
        </>
    );
};

export default Tasks;
