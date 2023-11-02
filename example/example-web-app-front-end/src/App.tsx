
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

import './App.css'
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';


import OrderForm from './OrderForm';
import { ItemController, OrdersList } from './OrderList';

function App() {
  let itemController:ItemController;
  return (
    <div>
      <Router>
      <Routes>
        <Route path="/" element={
        <div>
          <OrderForm></OrderForm>
          <OrdersList onInit={(ic)=>{
            itemController = ic
          }}></OrdersList>
        </div>} />
      </Routes>
    </Router>
    <ToastContainer />
    </div>
  );
}

export default App
