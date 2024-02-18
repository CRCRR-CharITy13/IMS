import "./App.css";

import { MantineProvider, ColorSchemeProvider, ColorScheme, GlobalStyles, Global } from "@mantine/core";
import { Notifications } from "@mantine/notifications";
import { useLocalStorage } from '@mantine/hooks';

import { BrowserRouter, Routes, Route } from "react-router-dom";

import Dashboard from "./routes/Dashboard";
import Login from "./routes/Login";
import Register from "./routes/Register";
import DashboardRedirect from "./components/dashboard/Redirect";
import Landing from "./routes/Landing";
import PageNotFound from "./routes/PageNotFound";

function App() {
    // const isDark = window.localStorage.getItem("dark") === "true";

    const [colorScheme, setColorScheme] = useLocalStorage<ColorScheme>({
        key: 'mantine-color-scheme',
        defaultValue: 'light',
        getInitialValueInEffect: true,
    });
    
    document.documentElement.setAttribute('data-theme', colorScheme)


    const toggleColorScheme = (value?: ColorScheme) => {
        setColorScheme(value || (colorScheme === 'dark' ? 'light' : 'dark'));
        document.documentElement.setAttribute('data-theme', colorScheme)
    }

    return (
        <ColorSchemeProvider colorScheme={colorScheme} toggleColorScheme={toggleColorScheme}>
            <MantineProvider 
                theme={{
                    colorScheme, primaryColor: 'teal'}}
                withGlobalStyles
            >
                <Notifications />
                <BrowserRouter>
                    <Routes>
                        <Route path="/" element={<Login />} />
                        <Route
                            path="/dashboard"
                            element={<DashboardRedirect />}
                        />
                        <Route
                            path="/dashboard/:handle"
                            element={<Dashboard />}
                        />
                        <Route path="/login" element={<Login />} />
                        <Route path="/register" element={<Register />} />
                        <Route path="*" element = {<PageNotFound />} />
                    </Routes>
                </BrowserRouter>
            </MantineProvider>
        </ColorSchemeProvider>
    );
}

export default App;
