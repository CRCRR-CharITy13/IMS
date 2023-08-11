import { useEffect, useState } from "react";
import { useNavigate, Link } from "react-router-dom";
const PageNotFound = () => {
    return (
        <div>
            <h1>404 - Page Not Found</h1>
            <Link to="/">Click here to return to the main page </Link>
            
        </div>
    );
};
export default PageNotFound;