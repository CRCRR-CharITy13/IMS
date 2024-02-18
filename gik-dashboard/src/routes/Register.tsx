import {Anchor, Container, Paper, PasswordInput, Text, Title} from "@mantine/core";

import styles from "../styles/Auth.module.scss";

import { TextInput, Checkbox, Button, Group } from "@mantine/core";
import { showNotification, hideNotification } from "@mantine/notifications";

import { useForm } from "@mantine/form";
import { useEffect, useState } from "react";
import {Link, useNavigate} from "react-router-dom";

const Register = () => {
    const navigate = useNavigate();

    const [registrationEnabled, setRegistrationEnabled] =
        useState<boolean>(false);

    const form = useForm({
        initialValues: {
            name: "",
            username: "",
            password: "",
            confPassword: "",
            authCode: "",
            agree: false,
        },
    });

    const lookupSignupCode = async (code: string) => {
        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/auth/scode?code=${code}`
        );

        if (response.status === 200) {
            const data = await response.json();
            setRegistrationEnabled(true);
            form.setFieldValue("name", data.data);
        }
    };

    useEffect(() => {
        if (!form.values.authCode.length) return;

        lookupSignupCode(form.values.authCode);
    }, [form.values.authCode]);

    const doRegister = async () => {
        if (form.values.password !== form.values.confPassword) {
            showNotification({
                color: "red",
                title: "Password Mismatch",
                message: "Password and password confirmation.tsx do not match.",
            });
            return;
        }

        setRegistrationEnabled(false);

        showNotification({
            id: "registrationLoading",
            loading: true,
            title: "Registering...",
            message: "Please wait while we create your account.",
        });

        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/auth/register`,
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    password: form.values.password,
                    passwordConf: form.values.confPassword,
                    signupCode: form.values.authCode,
                    eula: form.values.agree,
                }),
            }
        );

        hideNotification("registrationLoading");

        setRegistrationEnabled(true);

        const data = await response.json();

        if (data.success) {
            showNotification({
                color: "green",
                title: "Registered",
                message: "Registration successful. Redirecting to login...",
            });

            setTimeout(() => {
                navigate("/login", { replace: true });
            }, 1000);

            return;
        }

        showNotification({
            color: "red",
            title: "Registration Failed",
            message: data.message,
        });
    };

    return (
        
        <>
            <div>
                <Container size={420} my={40}>
                <Title
                    align="center"
                    sx={(theme) => ({ fontFamily: `Greycliff CF, ${theme.fontFamily}`, fontWeight: 900 })}
                >
                    Register
                    
                </Title>
                <Text color="red" size="sm" align="center" mt={5}>
                    Do have an account?{' '}
                    <Anchor size="sm" component="button">
                        {/* Link to Register */}
                    <Link to ="/login">Login</Link>
                    </Anchor>
                </Text>

                    <Paper withBorder shadow="md" p={30} mt={30} radius="md">
                    <form
                        onSubmit={(e) => {
                            e.preventDefault();
                            doRegister();
                        }}
                    >
                        <TextInput
                            required
                            disabled
                            label="Username"
                            placeholder={
                                form.values.name || "Waiting for auth code..."
                            }
                            {...form.getInputProps("username")}
                        />
                        <PasswordInput
                            required
                            label="Password"
                            mt="md" 
                            {...form.getInputProps("password")}
                        />
                        <PasswordInput
                            required
                            mt="md" 
                            label="Confirm Password"
                            {...form.getInputProps("confPassword")}
                        />
                        <TextInput
                            required
                            mt="md" 
                            label="Authorization Code"
                            placeholder="xxx"
                            {...form.getInputProps("authCode")}
                        />
                        <Checkbox
                            label="I agree to the terms and conditions and privacy policies."
                            sx={{
                                marginTop: "1rem",
                            }}
                            {...form.getInputProps("agree")}
                        />
                        <Group position="right" mt="md">
                            <Button
                                type="submit"
                                color="green"
                                disabled={!registrationEnabled}
                            >
                                Register
                            </Button>
                        </Group>
                    </form>
                    </Paper>

                </Container>
            </div>
        </>
    );
};

export default Register;
