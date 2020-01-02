import { Helmet } from 'react-helmet';
import * as React from 'react';
import { coFetchJSON } from '@console/internal/co-fetch';
import { Table, TableData, TableRow } from '@console/internal/components/factory';
import { sortable, SortByDirection } from '@patternfly/react-table';
import { PageHeading, ResourceLink } from '@console/internal/components/utils';
import { withStartGuide } from '../../../../public/components/start-guide';

interface PipelineRowListProps {
  obj: {
    name: string;
    namespace: string;
    info: {
      status: string;
      description: string;
    };
  };
  index: number;
  key?: string;
  style: object;
}

const PipelineRowList: React.FC<PipelineRowListProps> = ({ obj, index, key, style }) => {
  console.log('objs', obj, index, key);
  return (
    <TableRow id={index} index={index} trKey={key} style={style}>
      <TableData>{obj.name}</TableData>
      <TableData>
        <ResourceLink name={obj.namespace} kind="Namespace" />
      </TableData>
      <TableData>{obj.info.status}</TableData>
      <TableData>{obj.info.description}</TableData>
    </TableRow>
  );
};

/*
 // eslint-disable-next-line @typescript-eslint/no-unused-vars
 const PipelineRow: any = (args: any) => {
   console.log('helm data', args);
   return args.customData.map((d: any) => {
     return (
       <TableRow id={1} index={1} style={null} trKey="name" key="name">
         <TableData>{d.name}</TableData>
       </TableRow>
     );
   });
 };

 eslint-disable-next-line @typescript-eslint/no-unused-vars,react/prefer-stateless-function
*/
class HelmReleases extends React.Component {
  state = { response: null, loading: true };

  componentDidMount(): void {
    coFetchJSON('/api/console/helm/list')
      .then((resp) => {
        // eslint-disable-next-line react/no-unused-state
        this.setState({ response: resp, loading: false });
      })
      .catch((err) => console.error(err));
  }

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
    return (
      <span>
        <Helmet>
          <title>Helm Releases</title>
        </Helmet>
        <PageHeading title="Helm Releases">
          Select a project to view the list of helm releases
        </PageHeading>
        <hr className="odc-project-list-page__section-border" />
        <div style={{ padding: '20px' }}>
          {!this.state.loading && (
            <Table
              Header={() => [
                { title: 'Name', transform: [sortable], sortField: 'name' },
                { title: 'Namespace' },
                { title: 'Status' },
                { title: 'Description' },
              ]}
              aria-label="Nodes"
              data={this.state.response}
              Row={PipelineRowList}
              defaultSortOrder={SortByDirection.desc}
              loaded
              virtualize
            />
          )}
        </div>
      </span>
    );
  }
}

export default withStartGuide(HelmReleases);
