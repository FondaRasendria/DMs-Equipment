import React, {useEffect, useState} from "react";
import AppLayout from "../../layouts/AppLayout";
import { Link } from "react-router-dom";
import axios from "axios";

function ShowMonsters() {
    const [monsters, setMonsters] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch("http://localhost:8888/api/monster");

                if (!response.ok) {
                    throw new Error(`Failed to fetch data. Status: ${response.status}`);
                }

                const data = await response.json();
                setMonsters(data);
                console.log(data);

                const promises = data.map(async (monster) => {
                    const traitResponse = await fetch("http://localhost:8888/api/monstertrait/parentid/" + monster.id);
                    if(!traitResponse.ok) {
                        throw new Error(`Failed to fetch monstertraits for trait ${monster.id}. Status: Error`);
                    }
                    const traitData = await traitResponse.json();

                    const actionResponse = await fetch("http://localhost:8888/api/action/parentid/" + monster.id);
                    if(!actionResponse.ok) {
                        throw new Error(`Failed to fetch monstertraits for trait ${monster.id}. Status: Error`);
                    }
                    const actionData = await actionResponse.json();

                    return {...monster, traits: traitData.data, actions: actionData.data}
                });

                const updatedMonster = await Promise.all(promises);
                setMonsters(updatedMonster);
            } catch (error) {
                console.error("Error fetching user data:", error);
            }
        }
        fetchData();
    }, []);

    const abilityTable = (score) => {
        switch(score) {
            case 1:
                return -5;
            case 2 || 3:
                return -4;
            case 4 || 5:
                return -3;
            case 6 || 7:
                return -2;
            case 8 || 9:
                return -1;
            case 10 || 11:
                return 0;
            case 12 || 13:
                return 1;
            case 14 || 15:
                return 2;
            case 16 || 17:
                return 3;
            case 18 || 19:
                return 4;
            case 20 || 21:
                return 5;
            case 22 || 23:
                return 6;
            case 24 || 25:
                return 7;
            case 26 || 27:
                return 8;
            case 28 || 29:
                return 9;
            case 30:
                return 10;
        }
    }

    return(
        <AppLayout>
            <div className="mx-6">
            <button className="mb-2 items-center text-gray-500 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm px-3 py-1.5 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700" type="button">
                <a href="/addmonsters">
                    Add Monster
                </a>
            </button>
            
            {monsters.map((monster) => (
                <div key={monster.id} className="max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700">
                    <h5 className="text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{monster.name}</h5>
                    <h6 className="text-sm mb-2 font-light tracking-tight text-gray-900 dark:text-white italic">{monster.type}, <span>{monster.alignment}</span></h6>
                    <hr className="h-1 bg-red-700 border-0"></hr>
                    <p className="font-medium text-gray-700 dark:text-white">Armor Class: <span className="font-light">{monster.ac}</span></p>
                    <p className="font-medium text-gray-700 dark:text-white">HP: <span className="font-light">{monster.fixed_hp} <span>({monster.dice_hp})</span></span></p>
                    <p className="mb-1 font-medium text-gray-700 dark:text-white">Speed: <span className="font-light">{monster.speed}</span></p>
                    <hr className="mb-1 h-1 bg-red-700 border-0"></hr>
                    <p className="font-medium text-gray-700 dark:text-white">STR: <span className="font-light">{monster.str} ({abilityTable(monster.str)})</span></p>
                    <p className="font-medium text-gray-700 dark:text-white">DEX: <span className="font-light">{monster.dex} ({abilityTable(monster.dex)})</span></p>
                    <p className="font-medium text-gray-700 dark:text-white">CON: <span className="font-light">{monster.con} ({abilityTable(monster.con)})</span></p>
                    <p className="font-medium text-gray-700 dark:text-white">INT: <span className="font-light">{monster.int} ({abilityTable(monster.int)})</span></p>
                    <p className="font-medium text-gray-700 dark:text-white">WIS: <span className="font-light">{monster.wis} ({abilityTable(monster.wis)})</span></p>
                    <p className="mb-1 font-medium text-gray-700 dark:text-white">CHA: <span className="font-light">{monster.cha} ({abilityTable(monster.cha)})</span></p>
                    <hr className="mb-1 h-1 bg-red-700 border-0"></hr>
                    
                    {monster.saving_throws && (
                        <p className="font-medium text-gray-700 dark:text-white">Saving Throws: <span className="font-light">{monster.saving_throws}</span></p>
                    )}
                    {monster.skills && (
                        <p className="font-medium text-gray-700 dark:text-white">Skills: <span className="font-light">{monster.skills}</span></p>
                    )}
                    {monster.damage_immunities && (
                        <p className="font-medium text-gray-700 dark:text-white">Damage Immunities: <span className="font-light">{monster.damage_immunities}</span></p>
                    )}
                    {monster.damage_vulnerabilities && (
                        <p className="font-medium text-gray-700 dark:text-white">Damage Vulnerabilities: <span className="font-light">{monster.damage_vulnerabilities}</span></p>
                    )}
                    {monster.damage_resistences && (
                        <p className="font-medium text-gray-700 dark:text-white">Damage Resistences: <span className="font-light">{monster.damage_resistences}</span></p>
                    )}
                    {monster.condition_immunities && (
                        <p className="font-medium text-gray-700 dark:text-white">Condition Immunities: <span className="font-light">{monster.condition_immunities}</span></p>
                    )}
                    {monster.senses && (
                        <p className="font-medium text-gray-700 dark:text-white">Senses: <span className="font-light">{monster.senses}</span></p>
                    )}
                    {monster.language && (
                        <p className="font-medium text-gray-700 dark:text-white">Language: <span className="font-light">{monster.language}</span></p>
                    )}
                    {monster.cr && (
                        <p className="font-medium text-gray-700 dark:text-white">Challenge: <span className="font-light">{monster.cr}</span></p>
                    )}

                    <hr className="mb-1 h-1 bg-red-700 border-0"></hr>
                    {monster.traits && monster.traits.map((trait) => (
                        <p key={trait.id} className="font-medium text-gray-700 dark:text-white">&#x2022; {trait.name}, <span className="font-light">{trait.description}</span></p>
                    ))}
                    <hr className="mb-1 h-1 bg-red-700 border-0"></hr>
                    <p className="text-lg font-medium text-gray-700 dark:text-white">Action</p>
                    {monster.actions && monster.actions.map((action) => (
                        <p key={action.id} className="font-medium text-gray-700 dark:text-white">&#x2022; {action.name}, <span className="font-light">{action.description}</span></p>
                    ))}
                </div>
            ))}
            </div>
        </AppLayout>
    )
}
export default ShowMonsters;