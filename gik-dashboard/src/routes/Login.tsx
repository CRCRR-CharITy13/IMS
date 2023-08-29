import {
    Container,
    Modal,
    PasswordInput,
    Space,
    Checkbox,
    Text,
    Paper,
    Anchor,
    Title,
} from "@mantine/core";

import { TextInput, Button, Group } from "@mantine/core";
import { useForm } from "@mantine/form";

import { showNotification, hideNotification } from "@mantine/notifications";
import { useEffect, useState } from "react";
import { useNavigate, Link } from "react-router-dom";

const Login = () => {
    const navigate = useNavigate();

    const [loginEnabled, setLoginEnabled] = useState<boolean>(true);

    const form = useForm({
        initialValues: {
            username: "",
            password: "",
            rememberMe: false,
        },
    });

    const checkAuthStatus = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/auth/status`,
            {
                credentials: "include",
            }
        );

        if (response.status === 200) navigate("/dashboard", { replace: true });
    };
/*
    const do2FACheck = async (verification: string) => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/auth/tfa?verification=${verification}&username=${form.values.username}`
        );

        const data: {
            success: boolean;
            message: string;
            data: boolean;
        } = await response.json();

        if (!data.success) {
            showNotification({
                color: "red",
                title: "2FA Check Failed",
                message: data.message,
            });
            return;
        }

        if (data.data) {
            setTopVerification(verification);
            setAskTfa(true);
        } else {
            await doLogin(verification);
        }
    };*/

    const doPrelogin = async () => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/auth/prelogin`,
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    username: form.values.username,
                    password: form.values.password,
                }),
            }
        );

        const data: {
            success: boolean;
            message: string;
            data: string;
        } = await response.json();

        if (!data.success) {
            showNotification({
                color: "red",
                title: "Login Failed",
                message: data.message,
            });
            return;
        }
        await doLogin(data.data);
        //await do2FACheck(data.data);
    };

    const doLogin = async (verification: string) => {
        setLoginEnabled(false);

        showNotification({
            id: "loginLoading",
            loading: true,
            title: "Logging you in...",
            message: "Please wait while we authenticate you.",
        });

        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/auth/login?remember=${form.values.rememberMe}`,
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    username: form.values.username,
                    password: form.values.password,
                    verificationJWT: verification,
                }),
                credentials: "include",
            }
        );

        hideNotification("loginLoading");

        if (response.status === 200) {
            showNotification({
                color: "green",
                title: "Logged in",
                message: "Registration successful. Redirecting to dashboard...",
            });

            setTimeout(() => {
                navigate("/dashboard", { replace: true });
            }, 1000);
        } else {
            const data = await response.json();

            showNotification({
                color: "red",
                title: "Login Failed",
                message: data.message,
            });
        }

        setLoginEnabled(true);
    };

    useEffect(() => {
        checkAuthStatus();
    }, []);

//     return (
//         <>
//             <div className={`${styles.wrapper} ${styles.login}`}>
//                 <Container
//                     sx={{
//                         backgroundColor: "var(--inverted-text)",
//                         padding: "2rem",
//                         borderRadius: "5px",
//                         width: "40rem",
//                     }}
//                 >
//                     <h1
//                         style={{
//                             marginTop: 0,
//                         }}
//                     >
//                         Login
//                     </h1>
//                     <form
//                         onSubmit={(e) => {
//                             e.preventDefault();
//                             doPrelogin();
//                         }}
//                     >
//                         <TextInput
//                             required
//                             label="Username"
//                             placeholder="amy"
//                             {...form.getInputProps("username")}
//                         />
//                         <PasswordInput
//                             required
//                             placeholder="Password"
//                             label="Password"
//                             {...form.getInputProps("password")}
//                         />
//                         <Checkbox
//                             label="Remember me"
//                             color="green"
//                             {...form.getInputProps("rememberMe")}
//                         />
//                         <Button
//                             type="submit"
//                             color="green"
//                             disabled={!loginEnabled}
//                         >
//                             Login
//                         </Button>
//                     </form>
//                     <Button component={Link} to="/register" compact variant="white">
//                         <Text color="blue" size={"xs"}>Not Registered?</Text>
//                     </Button>
//                 </Container>

//             </div>
//         </>
//     );
// };

// export default Login;
return (
    <div>
    <Container size={420} my={40}>
      <Title
        align="center"
        sx={(theme) => ({ fontFamily: `Greycliff CF, ${theme.fontFamily}`, fontWeight: 900 })}
      >
        Gifts in Kind 
        
      </Title>
      <Text color="red" size="sm" align="center" mt={5}>
        Do not have an account yet?{' '}
        <Anchor size="sm" component="button">
            {/* Link to Register */}
          <Link to ="/register">Create account</Link>
        </Anchor>
      </Text>

              
               
      <Paper withBorder shadow="md" p={30} mt={30} radius="md">
        <form
            onSubmit={(e) => {
                e.preventDefault();
                doPrelogin();
            }}
        >
            <TextInput 
                label="Username" 
                placeholder="Enter username" 
                required 
                {...form.getInputProps("username")} 
            />
            <PasswordInput
                label="Password" 
                placeholder="Your password" 
                required
                mt="md" 
                {...form.getInputProps("password")} 
            />
            <Group 
                position="apart" 
                mt="lg"
            >
            <Checkbox
                label="Remember me"
                {...form.getInputProps("rememberMe")} 
            />
            <Anchor 
                component="button"
                size="sm"
            >
                Forgot password?
            </Anchor>
            </Group>
            <Button
                fullWidth mt="xl"
                type="submit"
                disabled={!loginEnabled}
            >
                Login
            </Button>
        </form>
      </Paper>
    </Container>
    </div>
  );
};

export default Login;
