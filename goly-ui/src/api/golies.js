export const getAllGolies = () => { 
    return fetch("http://localhost:3000/goly").then( resp => resp.json()) 
}


export const createGolyService = (longURL) => { 

    // let bodyContent = JSON.stringify({
    //     "goly": "asd",
    //     "redirect": longURL,
    //     "random": true
    // });

    // console.log(bodyContent);
    
    // return fetch("http://localhost:3000/goly", { 
    //     method: 'POST',
    //     headers: {"Content-Type":"application/json"},
    //     body: JSON.stringify(bodyContent),
    // })

    let headersList = {
        "Content-Type": "application/json"
       }
       
       let bodyContent = JSON.stringify({
           "goly": "",
           "redirect": longURL,
           "random": true
       });
       
    return fetch("http://localhost:3000/goly", { 
        method: "POST",
        body: bodyContent,
        headers: headersList
    })
 }