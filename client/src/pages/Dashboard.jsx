import React, {useEffect, useState} from "react";
import AppLayout from "../layouts/AppLayout";

function Dashboard() {
    const [races, setRaces] = useState([]);
    
    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch("http://localhost:8888/api/race");

                if (!response.ok) {
                    throw new Error(`Failed to fetch data. Status: ${response.status}`);
                }

                const data = await response.json();
                setRaces(data.Data);
            } catch (error) {
                console.error("Error fetching user data:", error);
            }
        }
        fetchData();
    }, []);
    
    return (
        <AppLayout>
            <h1 className="dark:text-white">Hello World</h1>
        </AppLayout>
    )
}

export default Dashboard;