import React, { FunctionComponent, useEffect, useState } from 'react';

import { ShoppingCartIcon, StarIcon, StoreLogoIcon, TagIcon, UserIcon } from './Icons';

interface Product {
	Name: string
	Price: number
	InStock: boolean
	Rating: number
	ImageUrl: string
}

export const App: FunctionComponent = () => {
  const [productList, setProductList] = useState<Product[]>([]);

  useEffect(() => {
    const getProductList = async () => {
      const response = await fetch("/api/product/get-list");
      const data = await response.json();
      setProductList(data);
    };
    getProductList();
  }, []);

  return (
    <div className="flex flex-col bg-gray-200 h-3/4 px-4 rounded-xl">
      <header className="flex my-4 items-center">
        <StoreLogoIcon className="w-10 h-10 mr-2" />
        <span className="text-3xl">WorldWide Software Architecture Summit</span>
        <span className="flex-grow"></span>
        <ShoppingCartIcon className="w-6 h-6" />
        <UserIcon className="w-6 h-6" />
      </header>
      <main className="flex-grow">
        <div className="bg-gray-300 rounded-xl p-4 mb-4 flex">
          <div className="flex-grow">
            <h1 className="text-2xl">
              <span className="block text-gray-800">Online Shopping</span>
              <span className="block text-green-600">Buy to be happier 🥳</span>
            </h1>
            <p className="my-4 text-base text-gray-500 text-lg">
              Anim aute id magna aliqua ad ad non deserunt sunt. Qui irure qui lorem cupidatat commodo. Elit sunt amet fugiat veniam occaecat fugiat aliqua.
            </p>
            <span className="inline-block cursor-pointer rounded-md shadow px-4 py-2 border border-transparent text-lg font-medium rounded-md text-white bg-green-600 hover:bg-green-700">
              Get started
            </span>
          </div>
          <img className="rounded-xl shadow" src="https://picsum.photos/300/200?color&random=2954" />
        </div>
        <div className="flex items-center">
          <TagIcon className="w-6 h-6 mr-2" />
          <span className="text-xl">Discounted Products</span>
          <span className="flex-grow"></span>
          <span className="cursor-pointer">View all ▸</span>
        </div>
        <div className="grid grid-flow-col gap-x-4 grid-cols-4">
          {productList.slice(0, 4).map((product, index) => {
            return (
              <div key={index} className="flex flex-col items-center bg-white rounded-xl mt-12 shadow">
                <img className="w-20 h-20 -mt-10 rounded-full border-2 border-gray-400 shadow" src={ product.ImageUrl } />
                <div className="w-full p-3">
                  <span className="text-lg">{ product.Name }</span>
                  <span className="flex items-center">
                    <StarIcon className="w-4 h-4 mr-1" />
                    <span className="text-md">{ product.Rating } / 5</span>
                  </span>
                  <span className="text-lg">${ product.Price }</span>
                  <span className="text-lg">{ product.Rating }</span>
                </div>
                <button className="py-1 px-4 border border-gray-300 rounded-lg mb-4 shadow">Add +</button>
              </div>
            );
          })}
        </div>
      </main>
      <footer className="my-4">
      </footer>
    </div>
  );
};
