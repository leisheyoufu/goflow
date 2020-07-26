package api

import (
	"net/http"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/leisheyoufu/goflow/pkg/registry"
)

type NodeResource struct {
	// normally one would use DAO (data access object)
	nodes map[string]registry.Node
}

// WebService creates a new service that can handle REST requests for User resources.
func (n NodeResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/nodes").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	tags := []string{"nodes"}

	ws.Route(ws.GET("/").To(n.listAllNodes).
		// docs
		Doc("get all nodes").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]registry.Node{}).
		Returns(200, "OK", []registry.Node{}))

	ws.Route(ws.GET("/{node-id}").To(n.getNode).
		// docs
		Doc("get a node").
		Param(ws.PathParameter("node-id", "identifier of the node").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(registry.Node{}). // on the response
		Returns(200, "OK", registry.Node{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(n.createNode).
		// docs
		Doc("create a node").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(registry.Node{})) // from the request

	ws.Route(ws.DELETE("/{node-id}").To(n.deleteNode).
		// docs
		Doc("delete a node").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("node-id", "identifier of the node").DataType("string")))

	return ws
}

// GET http://localhost:8080/nodes
//
func (n *NodeResource) listAllNodes(request *restful.Request, response *restful.Response) {
	list := []registry.Node{}
	for _, each := range n.nodes {
		list = append(list, each)
	}
	response.WriteEntity(list)
}

// GET http://localhost:8080/nodes/1
//
func (n *NodeResource) getNode(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("node-id")
	usr := n.nodes[id]
	if len(usr.ID) == 0 {
		response.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		response.WriteEntity(usr)
	}
}

// PUT http://localhost:8080/nodes/1
// <User><Id>1</Id><Name>Melissa</Name></User>
//
func (n *NodeResource) createNode(request *restful.Request, response *restful.Response) {
	//usr := User{ID: request.PathParameter("user-id")}
	node := registry.Node{}
	err := request.ReadEntity(&node)
	if err == nil {
		n.nodes[node.ID] = node
		response.WriteHeaderAndEntity(http.StatusCreated, node)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

// DELETE http://localhost:8080/nodes/1
//
func (n *NodeResource) deleteNode(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("node-id")
	delete(n.nodes, id)
}
