const API_URL = '/api';
let authToken = localStorage.getItem('authToken');

// Init
document.addEventListener('DOMContentLoaded', () => {
    if (authToken) {
        document.getElementById('login-overlay').style.display = 'none';
        fetchLogs();
    }
});

async function login() {
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    try {
        const res = await fetch('/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, password })
        });

        if (res.ok) {
            const data = await res.json();
            authToken = data.token;
            localStorage.setItem('authToken', authToken);
            document.getElementById('login-overlay').style.display = 'none';
            fetchLogs();
        } else {
            alert('Login failed');
        }
    } catch (e) {
        console.error(e);
        alert('Error logging in');
    }
}

async function fetchLogs() {
    if (!authToken) return;

    try {
        const res = await fetch(`${API_URL}/logs`, {
            headers: { 'Authorization': `Bearer ${authToken}` }
        });

        if (res.ok) {
            const logs = await res.json();
            renderLogs(logs);
            updateStats(logs);
        } else if (res.status === 401) {
            // Token expired
            localStorage.removeItem('authToken');
            location.reload();
        }
    } catch (e) {
        console.error('Failed to fetch logs', e);
    }
}

function renderLogs(logs) {
    const tbody = document.getElementById('logs-body');
    tbody.innerHTML = '';

    // Sort logs desc
    logs.sort((a, b) => new Date(b.timestamp) - new Date(a.timestamp));

    logs.forEach(log => {
        const tr = document.createElement('tr');

        let aiContent = '';
        if (log.analysis && log.analysis.is_anomaly) {
            aiContent = `<div class="ai-hint"><i class="fa-solid fa-wand-magic-sparkles"></i> ${log.analysis.suggested_fix || 'Anomaly Detected'}</div>`;
        }

        tr.innerHTML = `
            <td>${new Date(log.timestamp).toLocaleTimeString()}</td>
            <td><span class="tag ${log.level}">${log.level}</span></td>
            <td>${log.service}</td>
            <td>
                <div>${log.message}</div>
                ${aiContent}
            </td>
            <td>${log.analysis && log.analysis.severity_score ? log.analysis.severity_score + '/10' : '-'}</td>
        `;
        tbody.appendChild(tr);
    });
}

function updateStats(logs) {
    document.getElementById('stats-total').innerText = logs.length;
    const anomalies = logs.filter(l => l.analysis && l.analysis.is_anomaly).length;
    document.getElementById('stats-anomalies').innerText = anomalies;
}

// Auto refresh
setInterval(fetchLogs, 5000);
