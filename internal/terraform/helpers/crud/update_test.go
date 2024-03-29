package crud

// TestCrudUpdateCrossResourceContamination test CRUD helper: Update
// Had some issues where the client is shared across resources causing a clusterfuck during api commit, this is currently not needed since we no longer use a ptr for the client
// func TestCrudUpdateCrossResourceContamination(t *testing.T) {
// ctx := tflogtest.RootLogger(context.Background(), os.Stdout)

// 	// When Mock API Server
// 	address := "localhost:50015"
// 	uri := "/configure"
// 	key := "test-key"
// 	srv := &http.Server{
// 		Addr: address,
// 	}

// 	eList := api.NewExchangeList()

// 	e1 := eList.Add()
// 	e1.Expect(uri, key, `[{"op":"delete","path":["firewall","group","port-group","TF-grp","description","pre-grp"]},{"op":"set","path":["firewall","group","port-group","TF-grp","description","post-grp"]}]`)
// 	e1.Response(200, `{"success": true, "data": null, "error": null}`)

// 	e2 := eList.Add()
// 	e2.Expect(uri, key, `[{"op":"delete","path":["firewall","ipv4","name","TF-fw","description","pre-fw"]},{"op":"set","path":["firewall","ipv4","name","TF-fw","description","post-fw"]}]`)
// 	e2.Response(200, `{"success": true, "data": null, "error": null}`)

// 	api.Server(srv, eList)

// 	// When API Client
// 	providerDataClient := client.NewClient(ctx, "http://"+address, key, "test-agent", true)

// 	// With
// 	groupState := &fwgroup.FirewallGroupPortGroup{
// 		SelfIdentifier:                        basetypes.NewStringValue("TF-grp"),
// 		LeafFirewallGroupPortGroupDescrIPtion: basetypes.NewStringValue("pre-grp"),
// 	}
// 	fwState := &fwname.FirewallIPvfourName{
// 		SelfIdentifier:                     basetypes.NewStringValue("TF-fw"),
// 		LeafFirewallIPvfourNameDescrIPtion: basetypes.NewStringValue("pre-fw"),
// 	}

// 	// Should
// 	groupPlan := &fwgroup.FirewallGroupPortGroup{
// 		SelfIdentifier:                        basetypes.NewStringValue("TF-grp"),
// 		LeafFirewallGroupPortGroupDescrIPtion: basetypes.NewStringValue("post-grp"),
// 	}
// 	fwPlan := &fwname.FirewallIPvfourName{
// 		SelfIdentifier:                     basetypes.NewStringValue("TF-fw"),
// 		LeafFirewallIPvfourNameDescrIPtion: basetypes.NewStringValue("post-fw"),
// 	}

// 	// Run
// 	var wg sync.WaitGroup

// 	client1 := providerDataClient
// 	client2 := providerDataClient

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()

// 		err := update(ctx, client1, groupState, groupPlan)
// 		if err != nil {
// 			t.Logf("[grp] update error: %v", err.Error())
// 			t.Fail()
// 		}
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()

// 		err := update(ctx, client2, fwState, fwPlan)
// 		if err != nil {
// 			t.Logf("[fw] compare failed: %v", err.Error())
// 			t.Fail()
// 		}
// 	}()

// 	wg.Wait()

// // Validate API calls
// if len(eList.Unmatched()) > 0 {
// 	t.Logf("Total matched exchanges: %d", len(eList.Matched()))
// 	t.Errorf("Total unmatched exchanges: %d", len(eList.Unmatched()))
// 	t.Errorf("Next expected exchange match:\n%s", eList.Unmatched()[0].Sexpect())
// 	t.Errorf("Received request:\n%s", eList.Failed())
// }
// }
