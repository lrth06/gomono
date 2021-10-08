import './App.css';
import axios from 'axios'
import {useEffect, useState} from 'react'
function App() {
  const [msg, setMsg]= useState()
  const pingServer = async()=>{
    try{
      const res = await axios.get("http://localhost:5000")
      if(res){
        console.log(res)
        setMsg(res.data)
      }

    }catch(e){
      console.log(e)
    }
  }
  useEffect(()=>{
    pingServer()
  },[])
    return (
    <div className="App">
      <header className="App-header">
        {msg? <div>{msg}</div>:<div>Server not found</div>}
      
      </header>
    </div>
  );
}

export default App;
