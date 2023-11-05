chrome.omnibox.onInputEntered.addListener((text) => {
    console.log('inputEntered: ' + text);

    const query = encodeURIComponent(text.trim());
    const lookupUrl = `https://0593-2a02-a212-92c1-d880-ada8-cfc-662c-53b1.ngrok.io/links/${query}`;
  
    fetch(lookupUrl)
      .then(response => response.json()) 
      .then(data => {
        if (data && data.url) {
          chrome.tabs.create({ url: data.url });
        } else {
          console.error('Server returned an invalid URL or no URL at all');
        }
      })
      .catch(error => console.error('Error resolving ShortLink:', error));
  });
  
