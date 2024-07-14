import React, {useEffect, useState} from "react";
import AppLayout from "../../layouts/AppLayout";
import { Link } from "react-router-dom";
import axios from "axios";

function ShowRaces() {
    const [races, setRaces] = useState([]);

    useEffect(() => {
        const axiosRace = async () => {
            try {
                const response = await axios.get("http://localhost:8888/api/race");
                setRaces(response.data);
                console.log(response.data);
            } catch(error) {
                console.error("Error: ", error)
            }
        }
        axiosRace();
    }, []);

    return (
        <AppLayout>
            <div className="relative overflow-x-auto flex-col flex items-center shadow-md sm:rounded-lg">

                <div className="max-w-xl grid grid-cols-2 pb-4">
                    <div className="flex items-center">
                        <button className="inline-flex items-center text-gray-500 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm px-3 py-1.5 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700" type="button">
                            <a href="/addraces">
                                Add Race
                            </a>
                        </button>
                        
                    </div>
                    <div className="relative flex items-center">
                        <div className="absolute inset-y-0 left-0 rtl:inset-r-0 rtl:right-0 flex items-center ps-3 pointer-events-none">
                            <svg className="w-5 h-5 text-gray-500 dark:text-gray-400" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clipRule="evenodd"></path></svg>
                        </div>
                        <input type="text" id="table-search" className="block p-2 ps-10 text-sm text-gray-900 border border-gray-300 rounded-lg w-80 bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Search for items"/>
                    </div>
                </div>

                <table className="max-w-xl w-full text-sm text-center rtl:text-right text-gray-500 dark:text-gray-400">
                    <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                        <tr>
                            <th scope="col" className="px-6 py-3">
                                Race Name
                            </th>
                            <th scope="col" className="px-6 py-3">
                                Details
                            </th>
                        </tr>
                    </thead>
                    <tbody>
                        {races.map((race) => (
                            <tr key={race.id} className="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                                <th className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                                    {race.name}
                                </th>
                                <th className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white underline">
                                    <Link to={`${race.name}`} className="hover:text-gray-400">
                                        click
                                    </Link>
                                </th>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        </AppLayout>
    )
}

export default ShowRaces;