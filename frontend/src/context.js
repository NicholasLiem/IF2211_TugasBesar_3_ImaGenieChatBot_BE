import React, { useState, useContext } from "react";

const AppContext = React.createContext();

const AppProvider = ({ children }) => {
  const [selectedId, setSelectedId] = useState(null);

  return(
    <AppContext.Provider
        value={{
            selectedId,
            setSelectedId,
        }}
    >
        {children}
    </AppContext.Provider>
  )
};
export const useGlobalContext = () => {
  return useContext(AppContext)
}

export { AppContext, AppProvider }
