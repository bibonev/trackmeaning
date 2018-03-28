import { createMuiTheme } from 'material-ui/styles';

import indigo from 'material-ui/colors/indigo';
import pink from 'material-ui/colors/pink';
import red from 'material-ui/colors/red';

const muiTheme = createMuiTheme({
    palette: {
        contrastThreshold: 3,
        tonalOffset: 0.2,
        primary: indigo,
        secondary: pink,
        error: {
            main: red[500],
        },
    }
});

export default muiTheme;