import {
    Box,
    Button,
    Group,
    Input,
    InputWrapper,
    PasswordInput,
    Space,
    Switch,
} from "@mantine/core";
import { useForm } from "@mantine/form";
import { hideNotification, showNotification } from "@mantine/notifications";
import { useState } from "react";

import { containerStyles } from "../../styles/container";

const Settings = () => {
    const form = useForm({
        initialValues: {
            oldPassword: "",
            newPassword: "",
            newPasswordConfirm: "",
        },
    });

    const doChange = async () => {
        const values = form.values;

        if (values.newPassword !== values.newPasswordConfirm) {
            showNotification({
                message: "New passwords do not match.",
                title: "Password Mismatch",
                color: "red",
            });
            return;
        }

        showNotification({
            id: "changing",
            title: "Changing Password...",
            message: "Please wait while we process your request.",
            loading: true,
        });

        const response = await fetch(
            `${process.env.REACT_APP_API_URL}/settings/password`,
            {
                method: "PATCH",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
                body: JSON.stringify(form.values),
            }
        );

        hideNotification("changing");

        if (response.status !== 200) {
            const data = await response.json();

            showNotification({
                message: data.message,
                title: "Password Change Failed",
                color: "red",
            });
            return;
        }

        showNotification({
            message: "Password changed successfully.",
            title: "Password Changed",
            color: "green",
        });
    };

    const DoLogout = async () => {
        console.log("LOGGING OUT")
        await fetch(`${process.env.REACT_APP_API_URL}/auth/logout`,
            {credentials: "include",});
        window.location.reload();
    }

    return (
        <>
            <Box sx={containerStyles}>
                <h1>Settings</h1>
                <h2>Manage account and other user preferences.</h2>
            </Box>
            <Box sx={containerStyles}>
                <h3>Dashboard Theme</h3>
                <Space h="xl" />
                <Group>
                    <Switch
                        checked={window.localStorage.getItem("dark") === "true"}
                        onChange={() => {
                            if (
                                window.localStorage.getItem("dark") === "true"
                            ) {
                                window.localStorage.setItem("dark", "false");
                            } else {
                                window.localStorage.setItem("dark", "true");
                            }

                            window.location.reload();
                        }}
                    />
                    Dark Mode
                </Group>
            </Box>
            <Box sx={containerStyles}>
                <h3>Change Password</h3>
                <Space h="xl" />
                <form
                    onSubmit={(e) => {
                        e.preventDefault();
                        doChange();
                    }}
                >
                    <PasswordInput
                        required
                        label="Old Password"
                        {...form.getInputProps("oldPassword")}
                    />
                    <PasswordInput
                        required
                        label="New Password"
                        {...form.getInputProps("newPassword")}
                    />
                    <PasswordInput
                        required
                        label="Confirm New Password"
                        {...form.getInputProps("newPasswordConfirm")}
                    />
                    <Space h="xl" />
                    <Group position="right">
                        <Button type="submit" color="green">
                            Change Password
                        </Button>
                    </Group>
                </form>
            </Box>
            <Box sx={containerStyles}>
                <Button color="red" onClick={DoLogout}>
                    Logout
                </Button>
            </Box>
        </>
    );
};

export default Settings;
