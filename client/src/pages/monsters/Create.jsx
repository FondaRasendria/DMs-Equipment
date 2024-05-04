import React, {useState} from "react";
import { useNavigate, Link } from "react-router-dom";
import AppLayout from "../../layouts/AppLayout";
import axios from "axios";

function CreateMonster() {
    const navigate = useNavigate();

    const [traits, setTraits] = useState([]);
    const [actions, setActions] = useState([]);

    const handleAddTrait = () => {
        setTraits([...traits, { name: "", description: "" }]);
        console.log(traits);
    };
    const handleAddAction = () => {
        setActions([...actions, { name: "", description: "" }]);
        console.log(actions);
    };

    const handleSubmit = async () => {
        try {
            let axiosConfig = {
                headers: {
                    "Content-Type": "application/json"
                },
                method: "POST",
                url: "http://localhost:8888/api/monster",
                data: {

                },
            };
            let response = await axios(axiosConfig);
        }
        catch (error) {
            console.error(error)
        }
    }

    return (
        <AppLayout>
            <section className="bg-white dark:bg-gray-900">
            <div className="py-8 px-4 mx-auto max-w-2xl lg:py-16">
                <h2 className="mb-4 text-xl font-bold text-gray-900 dark:text-white">Add a new monster</h2>
                <form action="#">
                    <div className="grid gap-4 sm:grid-cols-2 sm:gap-6">
                        <div className="sm:col-span-2">
                            <label htmlFor="name" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Monster Name</label>
                            <input type="text" name="name" id="name" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Type monster name" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="type" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Type</label>
                            <input type="text" name="type" id="type" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Monster type" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="alignment" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Alignment</label>
                            <input type="text" name="alignment" id="alignment" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Alignment" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="ac" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Armor Class</label>
                            <input type="number" name="ac" id="ac" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="0" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="fixedhp" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Fixed HP</label>
                            <input type="number" name="fixedhp" id="fixedhp" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="0" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="dicehp" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Dice HP</label>
                            <input type="text" name="dicehp" id="dicehp" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Dice HP" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="speed" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Speed</label>
                            <input type="text" name="speed" id="speed" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="0 ft" required=""/>
                        </div>

                        <hr className="h-1 bg-red-700 border-0 sm:col-span-2"></hr>
                        
                        <div className="w-full">
                            <label htmlFor="str" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Strength (STR)</label>
                            <input type="number" name="str" id="str" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="0" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="dex" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Dexterity (DEX)</label>
                            <input type="number" name="dex" id="dex" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="0" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="con" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Constitution (CON)</label>
                            <input type="number" name="con" id="con" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="0" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="int" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Intelligence (INT)</label>
                            <input type="number" name="int" id="int" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="0" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="wis" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Wisdom (WIS)</label>
                            <input type="number" name="wis" id="wis" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="0" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="cha" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Charisma (CHA)</label>
                            <input type="number" name="cha" id="cha" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="0" required=""/>
                        </div>

                        <hr className="h-1 bg-red-700 border-0 sm:col-span-2"></hr>

                        <div className="w-full">
                            <label htmlFor="savingthrows" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Saving Throws</label>
                            <input type="text" name="savingthrows" id="savingthrows" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Saving throws" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="skills" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Skills</label>
                            <input type="text" name="skills" id="skills" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Skills" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="damageimun" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Damage Immunities</label>
                            <input type="text" name="damageimun" id="damageimun" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Damage immunities" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="damagevulne" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Damage Vulnerabilities</label>
                            <input type="text" name="damagevulne" id="damagevulne" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Damage vulnerabilities" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="damageresis" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Damage Resistences</label>
                            <input type="text" name="damageresis" id="damageresis" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Damage resistences" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="conditionimun" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Condition Immunities</label>
                            <input type="text" name="conditionimun" id="conditionimun" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Condition immunities" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="senses" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Senses</label>
                            <input type="text" name="senses" id="senses" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Senses" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="language" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Language</label>
                            <input type="text" name="language" id="language" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Language" required=""/>
                        </div>
                        <div className="w-full">
                            <label htmlFor="cr" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Challenge Rating (CR)</label>
                            <input type="text" name="cr" id="cr" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Challenge rating" required=""/>
                        </div>

                        <hr className="h-1 bg-red-700 border-0 sm:col-span-2"></hr>
 
                        <div className="w-full sm:col-span-2">
                            <button
                                type="button"
                                onClick={handleAddTrait}
                                className="inline-flex items-center px-5 py-2.5 mt-4 sm:mt-6 text-sm font-medium text-center text-white bg-primary-700 rounded-lg focus:ring-4 focus:ring-primary-200 dark:focus:ring-primary-900 hover:bg-primary-800"
                            >
                                Add Trait ({traits.length})
                            </button>
                        </div>

                        {traits.map((trait, index) => (
                            <div key={index} className="flex flex-col gap-4">
                                <input
                                    type="text"
                                    name={`trait-name-${index}`}
                                    className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                    placeholder="Trait name"
                                    value={trait.name}
                                    onChange={(e) => {
                                        // Update the name of the trait
                                        const updatedTraits = [...traits];
                                        updatedTraits[index].name = e.target.value;
                                        setTraits(updatedTraits);
                                    }}
                                />
                                <textarea
                                    name={`trait-description-${index}`}
                                    rows="4"
                                    className="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                    placeholder="Trait description"
                                    value={trait.description}
                                    onChange={(e) => {
                                        // Update the description of the trait
                                        const updatedTraits = [...traits];
                                        updatedTraits[index].description = e.target.value;
                                        setTraits(updatedTraits);
                                    }}
                                ></textarea>
                            </div>
                        ))}

                        <hr className="h-1 bg-red-700 border-0 sm:col-span-2"></hr>

                        <div className="w-full sm:col-span-2">
                            <button
                                type="button"
                                onClick={handleAddAction}
                                className="inline-flex items-center px-5 py-2.5 mt-4 sm:mt-6 text-sm font-medium text-center text-white bg-primary-700 rounded-lg focus:ring-4 focus:ring-primary-200 dark:focus:ring-primary-900 hover:bg-primary-800"
                            >
                                Add Action ({actions.length})
                            </button>
                        </div>

                        {actions.map((action, index) => (
                            <div key={index} className="flex flex-col gap-4">
                                <input
                                    type="text"
                                    name={`action-name-${index}`}
                                    className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                    placeholder="Action name"
                                    value={action.name}
                                    onChange={(e) => {
                                        // Update the name of the trait
                                        const updatedActions = [...actions];
                                        updatedActions[index].name = e.target.value;
                                        setActions(updatedActions);
                                    }}
                                />
                                <textarea
                                    name={`action-description-${index}`}
                                    rows="4"
                                    className="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                    placeholder="Action description"
                                    value={action.description}
                                    onChange={(e) => {
                                        // Update the description of the trait
                                        const updatedActions = [...actions];
                                        updatedActions[index].description = e.target.value;
                                        setActions(updatedActions);
                                    }}
                                ></textarea>
                            </div>
                        ))}

                        <hr className="h-1 bg-red-700 border-0 sm:col-span-2"></hr>

                        <div className="sm:col-span-2">
                            <label htmlFor="description" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Description</label>
                            <textarea id="description" rows="8" className="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Your description here"></textarea>
                        </div>
                        <div className="w-full sm:col-span-2">
                            <label htmlFor="source" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Source</label>
                            <input type="text" name="source" id="source" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Book source" required=""/>
                        </div>
                    </div>
                    <button type="submit" className="inline-flex items-center px-5 py-2.5 mt-4 sm:mt-6 text-sm font-medium text-center text-white bg-primary-700 rounded-lg focus:ring-4 focus:ring-primary-200 dark:focus:ring-primary-900 hover:bg-primary-800">
                        Add Monster
                    </button>
                </form>
            </div>
            </section>
        </AppLayout>
    )
}

export default CreateMonster;