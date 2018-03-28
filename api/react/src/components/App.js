import React from 'react';
import PropTypes from 'prop-types';
import {Link} from 'react-router';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';

import * as appActions from '../actions/appAction';
import * as Strings from './common/Strings';
import materialTheme from './common/Theme.js';

import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider'
import { withStyles } from 'material-ui/styles';
import AppBar from 'material-ui/AppBar';
import Toolbar from 'material-ui/Toolbar';
import Tooltip from 'material-ui/Tooltip';
import Typography from 'material-ui/Typography';
import Dialog, {
    DialogActions,
    DialogContent,
    DialogContentText,
    DialogTitle,
  } from 'material-ui/Dialog';
import Button from 'material-ui/Button';
import Grid from 'material-ui/Grid';
import Input, { InputLabel } from 'material-ui/Input';
import { MenuItem } from 'material-ui/Menu';
import { FormControl, FormHelperText } from 'material-ui/Form';
import Select from 'material-ui/Select';

import axios from 'axios';

const styles = theme => ({
    root: {
       minHeight: '100%'
    },
    flex: {
        flex: 1,
    },
    footer: {
        borderTop: '1px solid #9E9E9E',
    },
    space: {
        marginRight: '40px',
        marginLeft: '40px',
    }
})

class App extends React.Component {
    constructor() {
        super();
        this.state = {
            hoverOpacity: 1,
            animation: "none",
            open: false,
            language: {
                value: "en",
                text: "English"
            },
            inputValue: ""
        };
        this.toggleHover = this.toggleHover.bind(this);
        this.startLoading = this.startLoading.bind(this);
        this.onUpload = this.onUpload.bind(this);
        this.handleClose = this.handleClose.bind(this);
        this.handleChange = this.handleChange.bind(this);
    }

    toggleHover () {
        let value = this.state.hoverOpacity == 1 ? 0.5 : 1;
        this.setState({hoverOpacity: value})
    }

    startLoading(e) {
        e.click();
    }

    onUpload(el) {
        if (el.target.files.length === 1) {
            this.setState({animation: "2s rotateRight infinite linear"})
            this.props.actions.getResponse(el.target.files[0], this.state.language.value);
        }
    }

    handleClose() {
        this.setState({ open: false, inputValue: "" });
    };

    handleChange(event) {
        this.setState({ language: {
            value: event.target.value,
            text: event.target.name
        } });
    };

    componentWillReceiveProps(nextProps) {
        if (nextProps.meaning && this.state.open === false) {
            this.setState({open: true,  animation: "none" })
        }
    }

    render() {
        let inputElement;
        const { classes } = this.props;
        const { hoverOpacity, animation } = this.state;

        return (
            <MuiThemeProvider theme={materialTheme}>
                <Grid 
                    container
                    className={classes.root}
                    alignItems={'center'}
                    direction={'column'}
                    justify={'center'}
                >
                    <Grid key={0} item>
                        <FormControl className={classes.formControl}>
                            <InputLabel htmlFor="age-simple">Language</InputLabel>
                            <Select
                                value={this.state.language.value}
                                onChange={this.handleChange}
                                inputProps={{
                                name: 'language',
                                id: 'language-simple',
                                }}
                            >
                                <MenuItem value={"en"}>English</MenuItem>
                                <MenuItem value={"fr"}>French</MenuItem>
                                <MenuItem value={"it"}>Italian</MenuItem>
                                <MenuItem value={"bg"}>Bulgarian</MenuItem>
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid key={1} item>
                        <input 
                            ref={input => inputElement = input}
                            value={this.state.inputValue}
                            type='file'  
                            accept=".wav" 
                            style={{display:"none"}} 
                            onChange={this.onUpload}>
                        </input>​
                        <Tooltip id="tooltip-top" title="Choose Track" placement="top">
                            <img 
                            style={{
                                opacity: hoverOpacity,
                                cursor: "pointer",
                                transformOrigin:"50% 50%",
                                animation: animation
                            }} 
                            onMouseEnter={() => this.toggleHover()} 
                            onMouseLeave={() => this.toggleHover()}
                            onClick={() => this.startLoading(inputElement)}
                            src={"https://is1-ssl.mzstatic.com/image/thumb/Purple128/v4/cc/e1/c0/cce1c0a8-e23b-b629-68a1-099cbcd6c665/AppIcon.png/1200x630bb.png"} height="500" width="500" alt=""/>
                        </Tooltip>
                    </Grid>
                    <Grid key={2} item>
                        {
                            this.props.meaning && 
                            <Dialog
                                open={this.state.open}
                                onClose={this.handleClose}
                                aria-labelledby="alert-dialog-title"
                                aria-describedby="alert-dialog-description"
                                >
                                <DialogTitle id="alert-dialog-title">{this.props.meaning['Noun']}</DialogTitle>
                                <DialogContent>
                                    <DialogContentText id="alert-dialog-description">
                                        {this.props.meaning.KeyLyricsPrefix + ": "}
                                        <b>
                                        {
                                            this.props.meaning.KeyLyrics.map((val, index) => {
                                                return val + ", ";
                                            })
                                        }
                                        </b>
                                    </DialogContentText>
                                </DialogContent>
                                <DialogActions>
                                    <Button onClick={this.handleClose} color="primary">
                                        x
                                    </Button>
                                </DialogActions>
                            </Dialog>
                        }
                    </Grid>  
                    <Grid key={3} item className={classes.footer}>
                        <Typography component="h2" className={classes.space}>
                            {Strings.FOOTER_TITLE}
                        </Typography>
                    </Grid>
                </Grid>
            </MuiThemeProvider>
        );
    }
}

App.propTypes = {
    actions: PropTypes.object.isRequired,
    classes: PropTypes.object.isRequired,
    meaning: PropTypes.object
};

App.contextTypes = {
    router: PropTypes.object,
    location: PropTypes.object
}

function mapStateToProps(state, ownProps) {
    return {meaning: state.app.meaning};
}

function mapDispatchToProps(dispatch) {
    return {
        actions: bindActionCreators(appActions, dispatch)
    };
}

export default connect(mapStateToProps, mapDispatchToProps)(withStyles(styles)(App));