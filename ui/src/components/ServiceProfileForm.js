import React, { Component } from 'react';

import Select from "react-select";

import NetworkServerStore from "../stores/NetworkServerStore";
import SessionStore from "../stores/SessionStore";


class ServiceProfileForm extends Component {
  static contextTypes = {
    router: React.PropTypes.object.isRequired
  };

  constructor() {
    super();

    this.state = {
      serviceProfile: {
        serviceProfile: {},
      },
      networkServers: [],
      update: false,
      isAdmin: false,
    };

    this.handleSubmit = this.handleSubmit.bind(this);
    this.onNetworkServerChange = this.onNetworkServerChange.bind(this);
  }

  componentDidMount() {
    if (SessionStore.isAdmin()) {
      NetworkServerStore.getAll(9999, 0, (totalCount, networkServers) => {
        this.setState({
          serviceProfile: this.props.serviceProfile,
          networkServers: networkServers,
          isAdmin: true,
        });
      });
    } else {
      NetworkServerStore.getAllForOrganizationID(this.props.organizationID, 9999, 0, (totalCount, networkServers) => {
        this.setState({
          serviceProfile: this.props.serviceProfile,
          networkServers: networkServers,
          isAdmin: false,
        });
      });
    }

    SessionStore.on("change", () => {
      this.setState({
        isAdmin: SessionStore.isAdmin(),
      });
    });
  }

  componentWillReceiveProps(nextProps) {
    this.setState({
      serviceProfile: nextProps.serviceProfile,
      update: nextProps.serviceProfile.serviceProfile.serviceProfileID !== undefined,
    });
  }

  handleSubmit(e) {
    e.preventDefault();
    this.props.onSubmit(this.state.serviceProfile);
  }

  onChange(fieldLookup, e) {
    let lookup = fieldLookup.split(".");
    const fieldName = lookup[lookup.length-1];
    lookup.pop(); // remove last item

    let serviceProfile = this.state.serviceProfile;
    let obj = serviceProfile;

    for (const f of lookup) {
      obj = obj[f];
    }

    if (e.target.type === "number") {
      obj[fieldName] = parseInt(e.target.value, 10);
    } else if (e.target.type === "checkbox") {
      obj[fieldName] = e.target.checked;
    } else {
      obj[fieldName] = e.target.value;
    }

    this.setState({
      serviceProfile: serviceProfile,
    });
  }

  onNetworkServerChange(val) {
    let serviceProfile = this.state.serviceProfile;
    if (val != null) {
      serviceProfile.networkServerID = val.value;
    } else {
      serviceProfile.networkServerID = null;
    }
    this.setState({
      serviceProfile: serviceProfile,
    });
  }

  render() {
    const networkServerOptions = this.state.networkServers.map((networkServer, i) => {
      return {
        value: networkServer.id,
        label: networkServer.name,
      };
    });

    return(
      <form onSubmit={this.handleSubmit}>
        <fieldset disabled={!this.state.isAdmin}>
          <div className="form-group">
            <label className="control-label" htmlFor="name">Service-profile name</label>
            <input className="form-control" id="name" type="text" placeholder="e.g. my service-profile" required value={this.state.serviceProfile.name || ''} onChange={this.onChange.bind(this, 'name')} />
            <p className="help-block">
              A memorable name for the service-profile.
            </p>
          </div>
          <div className="form-group">
            <label className="control-label" htmlFor="networkServerID">Network-server</label>
            <Select
              name="networkServerID"
              options={networkServerOptions}
              value={this.state.serviceProfile.networkServerID}
              onChange={this.onNetworkServerChange}
              disabled={this.state.update}
            />
            <p className="help-block">
              The network-server on which this service-profile will be provisioned. After creating the service-profile, this value can't be changed.
            </p>
          </div>
          <div className="form-group">
            <label className="control-label" htmlFor="addGWMetadata">Add gateway meta-data</label>
            <div className="checkbox">
              <label>
                <input type="checkbox" name="addGWMetadata" id="addGWMetadata" checked={this.state.serviceProfile.serviceProfile.addGWMetadata} onChange={this.onChange.bind(this, 'serviceProfile.addGWMetadata')} /> Add gateway meta-data
              </label>
            </div>
            <p className="help-block">
              GW metadata (RSSI, SNR, GW geoloc., etc.) are added to the packet sent to the application-server.
            </p>
          </div>
          <div className="form-group">
            <label className="control-label" htmlFor="drMin">Minimum allowed data-rate</label>
            <input className="form-control" id="drMin" type="number" required value={this.state.serviceProfile.serviceProfile.drMin || 0} onChange={this.onChange.bind(this, 'serviceProfile.drMin')} />
            <p className="help-block">
              Minimum allowed data rate. Used for ADR.
            </p>
          </div>
          <div className="form-group">
            <label className="control-label" htmlFor="drMax">Maximum allowed data-rate</label>
            <input className="form-control" id="drMax" type="number" required value={this.state.serviceProfile.serviceProfile.drMax || 0} onChange={this.onChange.bind(this, 'serviceProfile.drMax')} />
            <p className="help-block">
              Maximum allowed data rate. Used for ADR.
            </p>
          </div>
        </fieldset>
        <hr />
        <div className={"btn-toolbar pull-right " + (this.state.isAdmin ? "" : "hidden")}>
          <a className="btn btn-default" onClick={this.context.router.goBack}>Go back</a>
          <button type="submit" className="btn btn-primary">Submit</button>
        </div>
      </form>
    );
  }
}

export default ServiceProfileForm;