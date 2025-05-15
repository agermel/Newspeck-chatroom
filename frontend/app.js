document.addEventListener('DOMContentLoaded', function() {
    let username = '';
    const loginBtn = document.getElementById('login-btn');
    const usernameInput = document.getElementById('username-input');
    const chatContainer = document.getElementById('chat-container');
    const messageForm = document.getElementById('message-form');
    const messageInput = document.getElementById('message-input');
    const messagesContainer = document.getElementById('messages');

    // Handle login
    loginBtn.addEventListener('click', function() {
        username = usernameInput.value.trim();
        if (username) {
            chatContainer.style.display = 'block';
            usernameInput.disabled = true;
            loginBtn.disabled = true;
            
            // Display welcome message
            const welcomeMsg = document.createElement('div');
            welcomeMsg.className = 'system-message';
            welcomeMsg.textContent = `Citizen ${username}, welcome to Newspeak Chat. Big Brother is watching you.`;
            messagesContainer.appendChild(welcomeMsg);
        }
    });
    
    // Connect to WebSocket for broadcasts
    const socket = new WebSocket('wss://newspeak.thusdaykfcv50.top/ws');
    
    // Handle form submission
    messageForm.addEventListener('submit', async function(e) {
        e.preventDefault();
        
        const message = messageInput.value.trim();
        if (message) {
            // Disable input while processing
            messageInput.disabled = true;
            messageInput.value = "Processing by Thought Police...";
            
            try {
                console.log('Sending message to filter API:', message);
                // First send to filter API
                const filterResponse = await fetch('https://newspeak.thusdaykfcv50.top/api/message/filter', {
                    method: 'POST',
                    mode: 'cors',
                    credentials: 'same-origin',
                    headers: {
                        'Content-Type': 'application/json',
                        'Accept': 'application/json'
                    },
                    body: JSON.stringify({
                        message: message,
                        user_id: 'citizen'
                    })
                });
                
                if (!filterResponse.ok) {
                    throw new Error(`HTTP error! status: ${filterResponse.status}`);
                }
                
                const filterData = await filterResponse.json();
                console.log('Received filter response:', filterData);
                
                // Only broadcast if not high danger level
                if (filterData.danger_level !== 'high') {
                    socket.send(JSON.stringify({
                        original_content: filterData.original,
                        content: filterData.filtered,
                        danger_level: filterData.danger_level,
                        modified: filterData.original !== filterData.filtered,
                        timestamp: new Date().toISOString(),
                        triggers: filterData.triggers,
                        note: filterData.note
                    }));
                }
                
                // Display the filtered message immediately
                displayMessage({
                    username: username,
                    original_content: filterData.original,
                    content: filterData.filtered,
                    danger_level: filterData.danger_level,
                    modified: filterData.original !== filterData.filtered,
                    timestamp: new Date().toISOString()
                });
            } catch (error) {
                console.error('Filter error:', error);
                alert('Ministry of Truth processing failed!');
            } finally {
                messageInput.disabled = false;
                messageInput.value = "";
            }
        }
    });
    
    // Handle incoming broadcast messages
    socket.addEventListener('message', function(event) {
        try {
            if (typeof event.data === 'string' && event.data.trim()) {
                const response = JSON.parse(event.data);
                displayMessage(response);
            } else {
                console.warn('Received empty or non-string WebSocket message:', event.data);
            }
        } catch (error) {
            console.error('Failed to parse WebSocket message:', error, 'Data:', event.data);
            const errorMsg = document.createElement('div');
            errorMsg.className = 'system-message error';
            errorMsg.textContent = 'Ministry of Truth message processing error';
            messagesContainer.appendChild(errorMsg);
        }
    });
    
    // Display message function
    function displayMessage(response) {
        const messageElement = document.createElement('div');
        messageElement.className = 'message';
        
        if (response.danger_level === 'high') {
            messageElement.innerHTML = `
                <div class="thoughtcrime">YOU COMMITTED THOUGHTCRIME</div>
                <div class="message-meta">${new Date(response.timestamp).toLocaleString()}</div>
            `;
        } else if (response.danger_level === 'none') {
            messageElement.classList.add('safe-message');
            messageElement.innerHTML = `
                <div class="message-header">Citizen ${response.username || username || 'Anonymous'}</div>
                <div class="message-content">${response.content || ''}</div>
                <span class="badge badge-safe">Approved by the Party</span>
                <div class="message-meta">
                    ${response.timestamp ? new Date(response.timestamp).toLocaleString() : new Date().toLocaleString()}
                </div>
            `;
        } else {
            if (response.modified) {
                messageElement.classList.add('modified-message');
                messageElement.innerHTML = `
                    <div class="message-header">Citizen ${response.username || username || 'Anonymous'}</div>
                    <div class="message-content">${response.content || ''}</div>
                    <span class="badge badge-modified">Message modified by the Party</span>
                    <div class="message-meta">
                        Original: "${response.original_content || ''}" | 
                        ${response.timestamp ? new Date(response.timestamp).toLocaleString() : new Date().toLocaleString()}
                    </div>
                `;
            } else {
                messageElement.innerHTML = `
                    <div class="message-header">Citizen ${response.username || username || 'Anonymous'}</div>
                    <div class="message-content">${response.content || ''}</div>
                    <div class="message-meta">${response.timestamp ? new Date(response.timestamp).toLocaleString() : new Date().toLocaleString()}</div>
                `;
            }
        }
        
        messagesContainer.appendChild(messageElement);
        messagesContainer.scrollTop = messagesContainer.scrollHeight;
    }
    
    // Handle connection errors
    socket.addEventListener('error', function(error) {
        console.error('WebSocket error:', error);
        alert('Connection to Ministry of Truth failed!');
    });
});
