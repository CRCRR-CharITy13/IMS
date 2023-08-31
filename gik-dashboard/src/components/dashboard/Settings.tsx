import {
    Box,
    Button,
    Group,
    PasswordInput,
    Space,
    Switch,
    useMantineTheme
} from "@mantine/core";
import { useForm } from "@mantine/form";
import { hideNotification, showNotification } from "@mantine/notifications";

import { useMantineColorScheme, ActionIcon} from '@mantine/core';

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

    return (
        <>
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
                        <Button type="submit">
                            Change Password
                        </Button>
                    </Group>
                </form>
            </Box>
        </>
    );
};

export default Settings;
