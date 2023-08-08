import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

const DashboardRedirect = () => {
    const navigate = useNavigate();

    useEffect(() => {
        navigate("/dashboard/analytics", { replace: true });
    }, []);

    return (
        <>
            <h1>Redirecting you...</h1>
        </>
    );
};

export default DashboardRedirect;
