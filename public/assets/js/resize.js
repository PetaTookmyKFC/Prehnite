function HandleResize(element) {
    return function (e) {
        
        // debugger;
        let offset = (element.style.width).replace("px", "") - e.offsetX;
        if (offset > 0) {
            return
        }
        
        e.preventDefault()
        console.log(offset)
        let max = (getComputedStyle(document.body).width).replace("px", "") - 25;
    
        let mouseMove = (e) => {
            e.preventDefault()
            let size = e.clientX + offset;


            if ( size > max )
            {
                return;
            }
            element.style.width = e.clientX + offset + "px";

            

        }
        let mouseUp = (e) => {
            localStorage.setItem("PD_SidebarSize", element.style.width)
            document.removeEventListener("mousemove",mouseMove )
            document.removeEventListener("mouseup", mouseUp)
        }

        document.addEventListener("mousemove", mouseMove)
        document.addEventListener("mouseup", mouseUp)
    
    }
}

/**
 * 
 * @param {Element} element 
*/
function AddResize(element) {
    // Check if event is already registered to this element
    if (element.getAttribute("Registered") == true) {
        return
    }
    // Set width
    element.style.width = localStorage.getItem("PD_SidebarSize") ?localStorage.getItem("PD_SidebarSize"): "100"
    // Register the event and set the flag to true.
    element.setAttribute("Registered", true)
    element.addEventListener("mousedown", HandleResize(element))


}


globalThis.InitDatabaseExplorer = () => {
    let item = document.getElementById("DatabaseWindow");
    console.log(item)
    if (item != null) {
        AddResize(item);
    }
    //     console.log("Checking!")
//     globalThis.RBars = Array.from(
//         document.getElementsByClassName("Resize")
//     )
//     for (i = 0; i < RBars.length; i++) {
//         console.log("Loaded: ", RBars[i]);
//         AddResize(RBars[i])
//     }
}
