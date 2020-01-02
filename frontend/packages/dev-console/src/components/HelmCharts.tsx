import { Helmet } from 'react-helmet';
import * as React from 'react';
import {
  Card,
  CardActions,
  CardBody,
  CardHead,
  CardHeader,
  Dropdown,
  DropdownItem,
  Gallery,
  GalleryItem,
  KebabToggle,
} from '@patternfly/react-core';
import classNames from 'classnames';
import { PageHeading } from '@console/internal/components/utils';
import { coFetchJSON } from '@console/internal/co-fetch';
import { withStartGuide } from '../../../../public/components/start-guide';

class HelmCharts extends React.Component {
  state = {
    isOpen: false,
  };

  onToggle = (isOpen) => {
    this.setState({
      isOpen,
    });
  };

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  onSelect = (event) => {
    this.setState({
      // eslint-disable-next-line react/no-access-state-in-setstate
      isOpen: !this.state.isOpen,
    });
  };

  onClick = (checked, event) => {
    const { target } = event;
    const value = target.type === 'checkbox' ? target.checked : target.value;
    const { name } = target;
    this.setState({ [name]: value });
  };

  render():
    | React.ReactElement<any, string | React.JSXElementConstructor<any>>
    | string
    | number
    | {}
    | React.ReactNodeArray
    | React.ReactPortal
    | boolean
    | null
    | undefined {
    const { isOpen } = this.state;

    const dropdownItems = [
      <DropdownItem
        key="install"
        component="button"
        onClick={() => {
          coFetchJSON
            .post(
              '/api/console/helm/install?ns=akash-helm-server&name=openshift-helm-demo-chart&url=https://technosophos.github.io/tscharts/mink-0.1.0.tgz',
              null,
              null,
            )
            .then(() => {
              alert('Chart Installed');
            })
            .catch((err: any) => console.error(err));
        }}
      >
        Install Chart
      </DropdownItem>,
    ];

    return (
      <span>
        <Helmet>
          <title>Helm Releases</title>
        </Helmet>
        <PageHeading title="Helm Catalog">Install chart from Helm Catalog</PageHeading>
        <hr className="odc-project-list-page__section-border" />
        <Gallery gutter="md" style={{ padding: '20px' }}>
          <GalleryItem>
            <Card className={classNames('catalog-tile-pf')} isHoverable>
              <CardHead>
                <CardActions>
                  <Dropdown
                    onSelect={this.onSelect}
                    toggle={<KebabToggle onToggle={this.onToggle} />}
                    isOpen={isOpen}
                    isPlain
                    dropdownItems={dropdownItems}
                    position="right"
                  />
                </CardActions>
              </CardHead>
              <CardHeader>Mink-v0.1.0</CardHeader>
              <CardBody>
                <p>PHP 5.3+ web browser emulator abstraction </p>
              </CardBody>
            </Card>
          </GalleryItem>
        </Gallery>
      </span>
    );
  }
}

export default withStartGuide(HelmCharts);
