package payspace

import (
	"context"
	"net/http"
	"strings"
	"testing"
)

func TestMetadataService_Get(t *testing.T) {
	xmlBody := `<?xml version="1.0" encoding="utf-8"?>
<edmx:Edmx Version="4.0" xmlns:edmx="http://docs.oasis-open.org/odata/ns/edmx">
  <edmx:DataServices>
    <Schema Namespace="PaySpace" xmlns="http://docs.oasis-open.org/odata/ns/edm">
      <EntityType Name="Employee">
        <Key><PropertyRef Name="EmployeeId"/></Key>
        <Property Name="EmployeeId" Type="Edm.Int32" Nullable="false"/>
        <Property Name="FirstName" Type="Edm.String"/>
      </EntityType>
    </Schema>
  </edmx:DataServices>
</edmx:Edmx>`

	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/$metadata", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(xmlBody))
	})
	_, client := testServerAndClient(t, mux)

	metadata, resp, err := client.Metadata.Get(context.Background(), 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if !strings.Contains(metadata, "edmx:Edmx") {
		t.Error("expected metadata to contain 'edmx:Edmx'")
	}
	if !strings.Contains(metadata, "EntityType Name=\"Employee\"") {
		t.Error("expected metadata to contain Employee EntityType")
	}
	if !strings.Contains(metadata, "FirstName") {
		t.Error("expected metadata to contain 'FirstName' property")
	}
}
