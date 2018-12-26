const btn1 = document.querySelector('.btn1');
const btn2 = document.querySelector('.btn2');
const customerDiv = document.getElementById('customer');
const customersDiv = document.getElementById('customers');

function loadCustomer(e) {
	const xhr = new XMLHttpRequest();

	xhr.open('GET', 'customer.json', true);

	xhr.onload = function() {
		if (this.status === 200) {
			// console.log(this.responseText);

			// turning json to an js object
			const customer = JSON.parse(this.responseText);

			const print = `
            
                <ul>
                    <li>ID: ${customer.id} </li>
                    <li>Name: ${customer.name} </li>
                    <li>City: ${customer.city} </li>
                    <li>phone-number: ${customer.phone} </li>
                </ul>
                    `;

			customerDiv.innerHTML = print;
		}
	};

	xhr.send();
}

function loadCustomers(e) {
	const xhr = new XMLHttpRequest();

	xhr.open('GET', 'customers.json', true);

	xhr.onload = function() {
		if (this.status === 200) {
			const customers = JSON.parse(this.responseText);

			let print = '';

			customers.forEach(function(customer) {
				print += `
            
                <ul>
                    <li>ID: ${customer.id} </li>
                    <li>name: ${customer.name} </li>
                    <li>city: ${customer.city} </li>
                    <li>phone-number: ${customer.phone} </li>
                        <br>
                </ul>
                    `;
			});

			customersDiv.innerHTML = print;
		}
	};
	xhr.send();
}

btn1.addEventListener('click', loadCustomer);
btn2.addEventListener('click', loadCustomers);

btn1.addEventListener('click', () => {
	customerDiv.classList.toggle('toggle');
});
btn2.addEventListener('click', () => {
	customersDiv.classList.toggle('toggle');
});
