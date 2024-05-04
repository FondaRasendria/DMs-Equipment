import React, { useEffect } from 'react'
import './App.css'
import {
  BrowserRouter as Router,
  Routes,
  Route,
  useNavigate,
} from 'react-router-dom'
import Dashboard from './pages/Dashboard'
import ShowRaces from './pages/races/Show'
import ShowRaceDetails from './pages/races/ShowDetails'
import CreateRace from './pages/races/Create'
import ShowMonsters from './pages/monsters/Show'
import CreateMonster from './pages/monsters/Create'

function RedirectToDashboard() {
  const navigate = useNavigate()

  useEffect(() => {
    navigate("/dashboard");
  }, [navigate]);

  return null;
}

function App() {

  return (
    <Router>
      <div className='App'>
        <Routes>
          <Route path='/dashboard' element={<Dashboard/>}/>
          <Route path='/races' element={<ShowRaces/>}/>
          <Route path='/races/:name' element={<ShowRaceDetails/>}/>
          <Route path='/addraces/' element={<CreateRace/>}/>
          <Route path='/monsters/' element={<ShowMonsters/>}/>
          <Route path='/addmonsters/' element={<CreateMonster/>}/>
          <Route path='*' element={<RedirectToDashboard/>}/>
        </Routes>
      </div>
    </Router>
  )
}

export default App
