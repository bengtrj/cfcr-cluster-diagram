var cy = cytoscape({
    container: document.getElementById('cy'),
    style: cytoscape.stylesheet()
        .selector('node')
        .css({
            'content': 'data(name)',
            'text-valign': 'center',
            'color': 'white',
            'text-outline-width': 5,
            'text-outline-color': '#888',
            'width': 80,
            'height': 80,
            'font-size': 17
        })
        .selector('edge')
        .css({
            'content': 'data(name)',
            'width': 8,
            'line-color': '#888',
            'target-arrow-color': '#888',
            'source-arrow-color': '#888',
            'target-arrow-shape': 'triangle'
        })
        .selector(':selected')
        .css({
            'background-color': 'black',
            'line-color': 'black',
            'target-arrow-color': 'black',
            'source-arrow-color': 'black',
            'text-outline-color': 'black'
        })
        .selector('$node > node')
        .css({
            'shape': 'roundrectangle',
            'text-valign': 'top',
            'height': 'auto',
            'width': 'auto',
            'background-color': '#ccc',
            'background-opacity': 0.333,
            'color': '#888',
            'text-outline-width':
                0,
            'font-size': 25
        })
        .selector('.master, #app')
        .css({
            'width': 120,
            'height': 120,
            'font-size': 25,
            'background-color': '#ccc',
            'background-opacity': 0.666,
        }).selector('.master-process')
        .css({
            'width': 80,
            'height': 80,
            'font-size': 15,
            'text-outline-width': 3,
            'background-color': '#ca4',
            'background-opacity': 0.666,
        }).selector('.master-etcd-cluster')
        .css({
            'font-size': 15,
            'text-outline-width': 0,
            'background-color': '#2a4',
            'background-opacity': 0.666,
        }).selector('.master-etcd')
        .css({
            'width': 70,
            'height': 70,
            'font-size': 15,
            'text-outline-width': 3,
            'background-color': '#2a4',
            'background-opacity': 0.666,
        })
        .selector('#worker-1 > node')
        .css({
            'shape': 'hexagon',
            'width': 50,
            'height': 50,
            'font-size': 10,
            'font-size': 10
        })
        .selector('#master, .ext')
        .css({
            'background-color': '#93CDDD',
            'text-outline-color': '#93CDDD',
            'line-color': '#93CDDD',
            'target-arrow-color': '#93CDDD'
        })
        .selector('#app, .app')
        .css({
            'background-color': '#F79646',
            'text-outline-color': '#F79646',
            'line-color': '#F79646',
            'target-arrow-color': '#F79646',
            'text-outline-color': '#F79646',
            'text-outline-width': 5,
            'color': '#fff'
        })
        .selector('#cy')
        .css({
            'background-opacity': 0,
            'border-width': 1,
            'border-color': '#aaa',
            'border-opacity': 0.5,
            'font-size': 50,
            'padding-top': 40,
            'padding-left': 40,
            'padding-bottom': 40,
            'padding-right': 40
        }),

    elements: {
        nodes: [
            {
                data: {
                    id: 'cy',
                    name: 'Clusters'
                }
            },

            {
                data: {
                    id: 'production',
                    name: 'production-mobile-backend',
                    parent: 'cy'
                },
            },

            {
                data: {
                    id: 'master-1',
                    name: 'Master 1',
                    parent: 'production',
                },
                classes: 'master'
            },

            {
                data: {
                    id: 'apiserver-master-1',
                    name: 'apiserver',
                    parent: 'master-1'
                },
                classes: 'master-process',
                position: {
                    x: 0,
                    y: 0
                }
            },

            {
                data: {
                    id: 'scheduler-master-1',
                    name: 'scheduler',
                    parent: 'master-1'
                },
                classes: 'master-process',
                position: {
                    x: 100,
                    y: 0
                }
            },

            {
                data: {
                    id: 'kube-controller-manager-master-1',
                    name: 'k-manager',
                    parent: 'master-1'
                },
                classes: 'master-process',
                position: {
                    x: 200,
                    y: 0
                }
            },

            {
                data: {
                    id: 'cloud-controller-manager-master-1',
                    name: 'c-manager',
                    parent: 'master-1'
                },
                classes: 'master-process',
                position: {
                    x: 300,
                    y: 0
                }
            },

            {
                data: {
                    id: 'etcd-cluster-master-1',
                    name: 'etcd-cluster',
                    parent: 'master-1'
                },
                classes: 'master-etcd-cluster'
            },

            {
                data: {
                    id: 'etcd-1-master-1',
                    name: 'etcd-1',
                    parent: 'etcd-cluster-master-1'
                },
                classes: 'master-etcd',
                position: {
                    x: 0,
                    y: 110
                }
            },

            {
                data: {
                    id: 'etcd-2-master-1',
                    name: 'etcd-2',
                    parent: 'etcd-cluster-master-1'
                },
                classes: 'master-etcd',
                position: {
                    x: 80,
                    y: 110
                }
            },

            {
                data: {
                    id: 'etcd-3-master-1',
                    name: 'etcd-3',
                    parent: 'etcd-cluster-master-1'
                },
                classes: 'master-etcd',
                position: {
                    x: 160,
                    y: 110
                }
            },

///////////////////////////////////



            {
                data: {
                    id: 'master-2',
                    name: 'Master 2',
                    parent: 'production',
                },
                classes: 'master'
            },

            {
                data: {
                    id: 'apiserver-master-2',
                    name: 'apiserver',
                    parent: 'master-2'
                },
                classes: 'master-process',
                position: {
                    x: 0,
                    y: 280
                }
            },

            {
                data: {
                    id: 'scheduler-master-2',
                    name: 'scheduler',
                    parent: 'master-2'
                },
                classes: 'master-process',
                position: {
                    x: 100,
                    y: 280
                }
            },

            {
                data: {
                    id: 'kube-controller-manager-master-2',
                    name: 'k-manager',
                    parent: 'master-2'
                },
                classes: 'master-process',
                position: {
                    x: 200,
                    y: 280
                }
            },

            {
                data: {
                    id: 'cloud-controller-manager-master-2',
                    name: 'c-manager',
                    parent: 'master-2'
                },
                classes: 'master-process',
                position: {
                    x: 300,
                    y: 280
                }
            },

            {
                data: {
                    id: 'etcd-cluster-master-2',
                    name: 'etcd-cluster',
                    parent: 'master-2'
                },
                classes: 'master-etcd-cluster'
            },

            {
                data: {
                    id: 'etcd-1-master-2',
                    name: 'etcd-1',
                    parent: 'etcd-cluster-master-2'
                },
                classes: 'master-etcd',
                position: {
                    x: 0,
                    y: 390
                }
            },

            {
                data: {
                    id: 'etcd-2-master-2',
                    name: 'etcd-2',
                    parent: 'etcd-cluster-master-2'
                },
                classes: 'master-etcd',
                position: {
                    x: 80,
                    y: 390
                }
            },

            {
                data: {
                    id: 'etcd-3-master-2',
                    name: 'etcd-3',
                    parent: 'etcd-cluster-master-2'
                },
                classes: 'master-etcd',
                position: {
                    x: 160,
                    y: 390
                }
            },


        ]
        // ,
        // edges: [
        //     {data: {source: 'apiserver', target: 'etcd-3'}},
        //     // {data: {source: 'apiserver', target: 'worker-1'}},
        //     {data: {source: 'apiserver', target: 'etcd-2'}},
        //     {data: {source: 'apiserver', target: 'etcd-1'}}
        // ]
    },

    layout: {
        name: 'preset'
    }
});

function addMasterNode(etcdInstances) {

}
