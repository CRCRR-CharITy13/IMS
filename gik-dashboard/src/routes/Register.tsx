import {Container, PasswordInput, Text} from "@mantine/core";

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
            <div className={`${styles.wrapper} ${styles.register}`}>
                <Container
                    sx={{
                        backgroundColor: "var(--inverted-text)",
                        padding: "2rem",
                        borderRadius: "5px",
                        width: "40rem",
                    }}
                >
                    <h1
                        style={{
                            marginTop: 0,
                        }}
                    >
                        Register
                    </h1>
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
                            {...form.getInputProps("password")}
                        />
                        <PasswordInput
                            required
                            label="Confirm Password"
                            {...form.getInputProps("confPassword")}
                        />
                        <TextInput
                            required
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
                        <Group position="apart" mt="md">
                            <Button component={Link} to="/login" compact variant="white">
                                <Text color="blue" size={"xs"}>Already Registered?</Text>
                            </Button>
                            <Button
                                type="submit"
                                color="green"
                                disabled={!registrationEnabled}
                            >
                                Register
                            </Button>
                        </Group>
                    </form>

                </Container>
            </div>
        </>
    );
};

export default Register;
