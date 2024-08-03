
import { Link } from "react-router-dom";

const TaskCard = ({ task }) => {

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

            <Link key={task.id} to={`/tasks/${task.id}`} className="text-decoration-none">
                <div className="card mb-3" style={{ transition: 'transform 0.2s', boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)' }}
                    onMouseOver={e => e.currentTarget.style.transform = 'translateY(-5px)'}
                    onMouseOut={e => e.currentTarget.style.transform = 'translateY(0)'}>
                    <div className="card-body">
                        <h5 className="card-title">{task.title}</h5>
                        <p className="card-text">{task.description}</p>
                        <span className={`badge ${getStatusBadgeClass(task.status)}`}>
                            Status: {task.status}
                        </span>
                    </div>
                </div>
            </Link>

        </>
    )
}

export default TaskCard;