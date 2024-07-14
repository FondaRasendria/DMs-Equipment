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
                setRace(data.data[0]);
            } catch (error) {
                console.error("Error fetching user data:", error);
            }
        }
        fetchData();
    }, []);

    useEffect(() => {
        const fetchData = async () => {
            try {
                if(race.id != undefined) {
                    const subraceResponse = await fetch("http://localhost:8888/api/subrace/parentid/" + race.id);

                    if (!subraceResponse.ok) {
                        throw new Error(`Failed to fetch data. Status: Error`);
                    }
                    const subraceData = await subraceResponse.json();
                    setSubraces(subraceData.data);

                    const promises = subraceData.data.map(async (subrace) => {
                        const traitResponse = await fetch("http://localhost:8888/api/trait/parentid/" + subrace.id);
                        if (!traitResponse.ok) {
                            throw new Error(`Failed to fetch traits for subrace ${subrace.id}. Status: Error`);
                        }
                        const traitData = await traitResponse.json();
                        
                        const traitPromises = traitData.data.map(async (trait) => {
                            const subtraitResponse = await fetch("http://localhost:8888/api/subtrait/parentid/" + trait.id);
                            if (!subtraitResponse.ok) {
                                throw new Error(`Failed to fetch subtraits for trait ${trait.id}. Status: Error`);
                            }
                            const subtraitData = await subtraitResponse.json();
                            return { ...trait, subtraits: subtraitData.data };
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
            <h1 className="dark:text-white text-4xl font-semibold py-3">{race.name}</h1>
            {subraces.map((subrace) => (
                <div key={subrace.id} className="py-4">
                    <h1 className="dark:text-white font-medium text-3xl">{subrace.source}</h1>
                    <h1 className="dark:text-white font-normal py-1 italic">{subrace.description}</h1>
                    <h1 className="dark:text-gray-400 font-medium py-1">&#x2022; Ability Score Increase. <span className="dark: text-white font-normal">{subrace.ability}</span></h1>
                    <h1 className="dark:text-gray-400 font-medium py-1">&#x2022; Age. <span className="dark:text-white font-normal">{subrace.age}</span></h1>
                    <h1 className="dark:text-gray-400 font-medium py-1">&#x2022; Alignment. <span className="dark:text-white font-normal">{subrace.alignment}</span></h1>
                    <h1 className="dark:text-gray-400 font-medium py-1">&#x2022; Size. <span className="dark:text-white font-normal">{subrace.size}</span></h1>
                    <h1 className="dark:text-gray-400 font-medium py-1">&#x2022; Speed. <span className="dark:text-white font-normal">{subrace.speed}</span></h1>
                    <h1 className="dark:text-gray-400 font-medium py-1">&#x2022; Language. <span className="dark:text-white font-normal">{subrace.language}</span></h1>

                    {subrace.traits && subrace.traits.map((trait) => (
                        <div key={trait.id}>
                            <h1 className="dark:text-gray-400 font-medium py-1">&#x2022; {trait.name}. <span className="dark:text-white font-normal">{trait.description}</span></h1>
                            {trait.subtraits && trait.subtraits.map((subtrait) => (
                                <div key={subtrait.id} className="mx-6">
                                    <h1 className="dark:text-white font-medium py-1">&#x2022; (Level {subtrait.level}) <span className="dark:text-white font-normal">{subtrait.description}</span></h1>
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