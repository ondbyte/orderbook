import React from 'react';
import 'tailwindcss/base.css'; // Import base styles
import 'tailwindcss/components.css'; // Import component styles
import 'tailwindcss/utilities.css'; // Import utility styles

import { toast } from 'react-toastify';
import { apiAddress } from './config';

interface ReservationState {
    side: string;
    quantity: string;
    price: string;
}

class Reservation extends React.Component<object, ReservationState> {
    constructor(props: object) {
        super(props);
        this.state = {
            side: 'sell',
            quantity: '',
            price: '',
        };
    }

    render() {
        return (
            <div className="max-w-xs mx-auto p-6 bg-white shadow-md rounded-md">
                <div className="mb-4">
                    <label className="block text-sm font-medium text-gray-600">Side:</label>
                    <select
                        className="mt-1 p-2 border rounded-md w-full"
                        value={this.state.side}
                        onChange={(c) => {
                            this.setState({
                                side: c.target.value,
                            });
                        }}
                    >
                        <option value="0">Sell</option>
                        <option value="1">Buy</option>
                    </select>
                </div>
                <div className="mb-4">
                    <label className="block text-sm font-medium text-gray-600">Quantity:</label>
                    <input
                        className="mt-1 p-2 border rounded-md w-full"
                        name="quantity"
                        type="number"
                        value={this.state.quantity}
                        onChange={(e) => {
                            this.setState({
                                quantity: e.target.value,
                            });
                        }}
                    />
                </div>
                <div className="mb-4">
                    <label className="block text-sm font-medium text-gray-600">Price:
                    <br />
                    <small>leave this empty if you are placing market order</small></label>
                    
                    <input
                        className="mt-1 p-2 border rounded-md w-full"
                        name="price"
                        type="number"
                        value={this.state.price}
                        onChange={(e) => {
                            this.setState({
                                price: e.target.value,
                            });
                        }}
                    />
                </div>
                <div>
                    <button onClick={()=>{
                            const api = apiAddress+"/order";
                            console.log(api)
                        if (this.state.side==""){
                            showToast("side is required")
                            return
                        }
                        if (this.state.quantity==""){
                            showToast("quantity is required")
                            return
                        }
                        if (this.state.price==""){
                            showToast("market order")
                        }
                    }} className="bg-transparent hover:bg-green-600 text-black font-semibold hover:text-white py-2 px-4 border border-black hover:border-transparent rounded">
                        Place order
                    </button>
                </div>
            </div>
        );
    }
}

function showToast(s:string) {
    toast.info(s, {
      position: 'top-right',
      autoClose: 3000, // Close the toast after 3 seconds (3000 milliseconds)
      hideProgressBar: false,
      closeOnClick: true,
      pauseOnHover: true,
      draggable: true,
      progress: undefined,
    });
  }

export default Reservation;
