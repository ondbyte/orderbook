import React, {  } from 'react';

type Order = {
  id :string
  quantity:number
  price:number|null
}

export interface ItemController {
  addItem(o:Order):void
}

type Props = {
  onInit: (controller: ItemController) => void;
};

interface OrdersState {
  orders:Order[]
}

export class OrdersList extends React.Component<Props, OrdersState> implements ItemController {
  constructor(props: Props) {
      super(props);
      this.state = {
        orders:[]
      };
      this.props.onInit(this)
  }
  addItem(o: Order): void {
    this.setState({
      orders:[...this.state.orders,o]
    })
  }

  render() {
  return (<table className="min-w-full bg-white border border-gray-200">
  <thead>
    <tr>
    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Sl.no</th>
      <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">id</th>
      <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">qty</th>
      <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">price</th>
    </tr>
  </thead>
  <tbody className="divide-y divide-gray-200">
    {
      this.state.orders.map((item,i)=>(
        <tr>
      <td className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{i}</td>
      <td className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{item.id}</td>
      <td className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{item.quantity}</td>
      <td className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">{item.price}</td>
    </tr>
      ))
    }
  </tbody>
</table>
);
  }
}

