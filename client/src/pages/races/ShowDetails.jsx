import React, {useEffect, useState} from "react";
import AppLayout from "../../layouts/AppLayout";
import { useParams, useNavigate, Link } from "react-router-dom";

function ShowRaceDetails() {
    const {name} = useParams();

    const [race, setRace] = useState([]);
    const [subraces, setSubraces] = useState([]);
    const [traits, setTraits] = useState([]);
    const [subtraits, setSubtraits] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch("http://localhost:8888/api/race/name/" + name);

                if (!response.ok) {
                    throw new Error(`Failed to fetch data. Status: ${response.status}`);
                }

                const data = await response.json();
                setRace(data.Data);
            } catch (error) {
                console.error("Error fetching user data:", error);
            }
        }
        fetchData();
    }, []);

    useEffect(() => {
        const fetchData = async () => {
            try {
                if(race.Id != undefined) {
                    const subraceResponse = await fetch("http://localhost:8888/api/subrace/parent/" + race.Id);

                    if (!subraceResponse.ok) {
                        throw new Error(`Failed to fetch data. Status: ${subraceResponse.status}`);
                    }
                    const subraceData = await subraceResponse.json();
                    setSubraces(subraceData.Data);

                    const promises = subraceData.Data.map(async (subrace) => {
                        const traitResponse = await fetch("http://localhost:8888/api/trait/parent/" + subrace.Id);
                        if (!traitResponse.ok) {
                            throw new Error(`Failed to fetch traits for subrace ${subrace.Id}. Status: ${traitResponse.Status}`);
                        }
                        const traitData = await traitResponse.json();
                        
                        const traitPromises = traitData.Data.map(async (trait) => {
                            const subtraitResponse = await fetch("http://localhost:8888/api/subtrait/parent/" + trait.Id);
                            if (!subtraitResponse.ok) {
                                throw new Error(`Failed to fetch subtraits for trait ${trait.Id}. Status: ${subtraitResponse.Status}`);
                            }
                            const subtraitData = await subtraitResponse.json();
                            return { ...trait, subtraits: subtraitData.Data };
                        });

                        const updatedTraits = await Promise.all(traitPromises);
                        return { ...subrace, traits: updatedTraits };
                    });

                    const updatedSubraces = await Promise.all(promises);
                    setSubraces(updatedSubraces);
                }
            } catch (error) {
                console.error("Error fetching user data:", error);
            }
        }
        fetchData();
    }, [race]);

    return(
        <AppLayout>
            <div className="mx-6">
            <h1 className="dark:text-white text-4xl font-semibold py-3">{race.Name}</h1>
            {subraces.map((subrace) => (
                <div key={subrace.Id} className="py-4">
                    <h1 className="dark:text-white font-medium text-3xl">{subrace.Source}</h1>
                    <h1 className="dark:text-white font-normal py-1 italic">{subrace.Description}</h1>
                    <h1 className="dark:text-gray-400 font-medium py-1">&#x2022; Ability Score Increase. <span className="dark: text-white font-normal">{subrace.Ability}</span></h1>
                    <h1 className="dark:text-gray-400 font-medium py-1">&#x2022; Age. <span className="dark:text-white font-normal">{subrace.Age}</span></h1>
                    <h1 className="dark:text-gray-400 font-medium py-1">&#x2022; Alignment. <span className="dark:text-white font-normal">{subrace.Alignment}</span></h1>
                    <h1 className="dark:text-gray-400 font-medium py-1">&#x2022; Size. <span className="dark:text-white font-normal">{subrace.Size}</span></h1>
                    <h1 className="dark:text-gray-400 font-medium py-1">&#x2022; Speed. <span className="dark:text-white font-normal">{subrace.Speed}</span></h1>
                    <h1 className="dark:text-gray-400 font-medium py-1">&#x2022; Language. <span className="dark:text-white font-normal">{subrace.Language}</span></h1>

                    {subrace.traits && subrace.traits.map((trait) => (
                        <div key={trait.Id}>
                            <h1 className="dark:text-gray-400 font-medium py-1">&#x2022; {trait.Name}. <span className="dark:text-white font-normal">{trait.Description}</span></h1>
                            {trait.subtraits && trait.subtraits.map((subtrait) => (
                                <div key={subtrait.Id} className="mx-6">
                                    <h1 className="dark:text-white font-medium py-1">&#x2022; (Level {subtrait.Level}) <span className="dark:text-white font-normal">{subtrait.Description}</span></h1>
                                </div>
                            ))}
                        </div>
                    ))}    
                </div>
            ))}
            </div>
        </AppLayout>
    )
}

export default ShowRaceDetails;