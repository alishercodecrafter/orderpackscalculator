// web/static/js/main.js
document.addEventListener('DOMContentLoaded', function() {
    refreshPackSizes();
});

function refreshPackSizes() {
    fetch('/api/packs')
        .then(response => response.json())
        .then(packs => {
            const packSizesList = document.getElementById('packSizesList');
            packSizesList.innerHTML = '';

            packs.forEach(pack => {
                const row = document.createElement('tr');
                row.innerHTML = `
                    <td>${pack.size}</td>
                    <td><button class="btn-delete" onclick="removePackSize(${pack.size})">Remove</button></td>
                `;
                packSizesList.appendChild(row);
            });
        })
        .catch(error => console.error('Error fetching pack sizes:', error));
}

function addPackSize() {
    const input = document.getElementById('newPackSize');
    const packSize = parseInt(input.value);

    if (!packSize || isNaN(packSize) || packSize <= 0) {
        alert('Please enter a valid pack size (greater than zero)');
        return;
    }

    fetch('/api/packs', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            pack: {
                size: packSize
            }
        })    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            alert(data.error);
        } else {
            input.value = '';
            refreshPackSizes();
        }
    })
    .catch(error => console.error('Error adding pack size:', error));
}

function removePackSize(size) {
    fetch(`/api/packs/${size}`, {
        method: 'DELETE'
    })
    .then(response => response.json())
    .then(data => {
        if (data.error) {
            alert(data.error);
        } else {
            refreshPackSizes();
        }
    })
    .catch(error => console.error('Error removing pack size:', error));
}

function calculatePacks() {
    const orderSizeInput = document.getElementById('orderSize');
    const orderSize = parseInt(orderSizeInput.value);

    if (!orderSize || isNaN(orderSize) || orderSize <= 0) {
        alert('Please enter a valid order size (greater than zero)');
        return;
    }

    fetch('/api/calculate', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ orderSize })
    })
    .then(response => response.json())
    .then(result => {
        if (result.error) {
            alert(result.error);
            return;
        }

        displayResults(result);
    })
    .catch(error => console.error('Error calculating packs:', error));
}

function displayResults(result) {
    const resultSection = document.getElementById('resultSection');
    const resultBody = document.getElementById('resultBody');

    resultBody.innerHTML = '';

    // Sort pack sizes in descending order for display
    const packSizes = Object.keys(result.packs).map(Number).sort((a, b) => b - a);

    packSizes.forEach(packSize => {
        const count = result.packs[packSize];
        const row = document.createElement('tr');
        row.innerHTML = `
            <td>${packSize}</td>
            <td>${count}</td>
        `;
        resultBody.appendChild(row);
    });

    resultSection.style.display = 'block';
}