import { BackgroundImage, Box, Button, Group } from "@mantine/core";
import { Link } from "react-router-dom";

const Landing = () => {
    return (
        <>
            <div
                style={{
                    height: "100%",
                    width: "100%",
                    backgroundImage: "url('/home_bg.jpg')",
                    backgroundSize: "cover",
                    backgroundAttachment: "fixed",
                    backgroundRepeat: "no-repeat",
                    backgroundBlendMode: "multiply",
                    backgroundColor: "#999999",
                    display: "flex",
                    justifyContent: "center",
                    alignItems: "center",
                }}
            >
                <Box
                    sx={{
                        padding: "1rem",
                        borderRadius: "10px",
                    }}
                >
                    <Group>
                        <Link to="/login">
                            <Button>Login</Button>
                        </Link>
                        <Link to="/register">
                            <Button color="green">Register</Button>
                        </Link>
                    </Group>
                </Box>
            </div>
        </>
    );
};

export default Landing;
