import React, {useEffect, useState} from "react";
import AppLayout from "../../layouts/AppLayout";
import { Link } from "react-router-dom";

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
                setMonsters(data.Data);

                const promises = data.Data.map(async (monster) => {
                    const traitResponse = await fetch("http://localhost:8888/api/monstertrait/parent/" + monster.Id);
                    if(!traitResponse.ok) {
                        throw new Error(`Failed to fetch monstertraits for trait ${monster.Id}. Status: ${traitResponse.Status}`);
                    }
                    const traitData = await traitResponse.json();

                    const actionResponse = await fetch("http://localhost:8888/api/action/parent/" + monster.Id);
                    if(!actionResponse.ok) {
                        throw new Error(`Failed to fetch monstertraits for trait ${monster.Id}. Status: ${actionResponse.Status}`);
                    }
                    const actionData = await actionResponse.json();

                    return {...monster, traits: traitData.Data, actions: actionData.Data}
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
                <div key={monster.Id} className="max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700">
                    <h5 className="text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{monster.Name}</h5>
                    <h6 className="text-sm mb-2 font-light tracking-tight text-gray-900 dark:text-white italic">{monster.Type}, <span>{monster.Alignment}</span></h6>
                    <hr className="h-1 bg-red-700 border-0"></hr>
                    <p className="font-medium text-gray-700 dark:text-white">Armor Class: <span className="font-light">{monster.AC}</span></p>
                    <p className="font-medium text-gray-700 dark:text-white">HP: <span className="font-light">{monster.FixedHP} <span>({monster.DiceHP})</span></span></p>
                    <p className="mb-1 font-medium text-gray-700 dark:text-white">Speed: <span className="font-light">{monster.Speed}</span></p>
                    <hr className="mb-1 h-1 bg-red-700 border-0"></hr>
                    <p className="font-medium text-gray-700 dark:text-white">STR: <span className="font-light">{monster.STR} ({abilityTable(monster.STR)})</span></p>
                    <p className="font-medium text-gray-700 dark:text-white">DEX: <span className="font-light">{monster.DEX} ({abilityTable(monster.DEX)})</span></p>
                    <p className="font-medium text-gray-700 dark:text-white">CON: <span className="font-light">{monster.CON} ({abilityTable(monster.CON)})</span></p>
                    <p className="font-medium text-gray-700 dark:text-white">INT: <span className="font-light">{monster.INT} ({abilityTable(monster.INT)})</span></p>
                    <p className="font-medium text-gray-700 dark:text-white">WIS: <span className="font-light">{monster.WIS} ({abilityTable(monster.WIS)})</span></p>
                    <p className="mb-1 font-medium text-gray-700 dark:text-white">CHA: <span className="font-light">{monster.CHA} ({abilityTable(monster.CHA)})</span></p>
                    <hr className="mb-1 h-1 bg-red-700 border-0"></hr>
                    
                    {monster.SavingThrows && (
                        <p className="font-medium text-gray-700 dark:text-white">Saving Throws: <span className="font-light">{monster.SavingThrows}</span></p>
                    )}
                    {monster.Skills && (
                        <p className="font-medium text-gray-700 dark:text-white">Skills: <span className="font-light">{monster.Skills}</span></p>
                    )}
                    {monster.DamageImmunities && (
                        <p className="font-medium text-gray-700 dark:text-white">Damage Immunities: <span className="font-light">{monster.DamageImmunities}</span></p>
                    )}
                    {monster.DamageVulnerabilities && (
                        <p className="font-medium text-gray-700 dark:text-white">Damage Vulnerabilities: <span className="font-light">{monster.DamageVulnerabilities}</span></p>
                    )}
                    {monster.DamageResistences && (
                        <p className="font-medium text-gray-700 dark:text-white">Damage Resistences: <span className="font-light">{monster.DamageResistences}</span></p>
                    )}
                    {monster.ConditionImmunities && (
                        <p className="font-medium text-gray-700 dark:text-white">Condition Immunities: <span className="font-light">{monster.ConditionImmunities}</span></p>
                    )}
                    {monster.Senses && (
                        <p className="font-medium text-gray-700 dark:text-white">Senses: <span className="font-light">{monster.Senses}</span></p>
                    )}
                    {monster.Language && (
                        <p className="font-medium text-gray-700 dark:text-white">Language: <span className="font-light">{monster.Language}</span></p>
                    )}
                    {monster.CR && (
                        <p className="font-medium text-gray-700 dark:text-white">Challenge: <span className="font-light">{monster.CR}</span></p>
                    )}

                    <hr className="mb-1 h-1 bg-red-700 border-0"></hr>
                    {monster.traits && monster.traits.map((trait) => (
                        <p key={trait.Id} className="font-medium text-gray-700 dark:text-white">&#x2022; {trait.Name}, <span className="font-light">{trait.Description}</span></p>
                    ))}
                    <hr className="mb-1 h-1 bg-red-700 border-0"></hr>
                    <p className="text-lg font-medium text-gray-700 dark:text-white">Action</p>
                    {monster.actions && monster.actions.map((action) => (
                        <p key={action.Id} className="font-medium text-gray-700 dark:text-white">&#x2022; {action.Name}, <span className="font-light">{action.Description}</span></p>
                    ))}
                </div>
            ))}
            </div>
        </AppLayout>
    )
}
export default ShowMonsters;