import { Button, Group, Modal, Text} from "@mantine/core";
import {Dispatch, SetStateAction} from "react";


export const ConfirmationModal = (
    {
        opened,
        setOpened,
        command,
        message,

    }: {
        opened: boolean;
        setOpened: Dispatch<SetStateAction<boolean>>;
        command: ()=>void;
        message: string;
    }) => {
    return (
        <>
            <Modal
                title={"Confirmation"}
                opened={opened}
                onClose={() => {
                    setOpened(false);
                }}
            >
                <Text>{message}</Text>
                <br/>
                <Group position={"right"}>
                    <Button color={"gray"} onClick={() => {setOpened(false);}}>Cancel</Button>
                    <Button color={"red"} onClick={() => {command(); setOpened(false);}}>Confirm</Button>
                </Group>
            </Modal>
        </>
    );
}
