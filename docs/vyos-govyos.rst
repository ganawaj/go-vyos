:lastproofread: 2024-03-10

.. _vyos-govyos:

go-vyos
======

go-vyos is a Go library designed for interacting with VyOS devices through
their API. This documentation is intended to guide you in using go-vyos for
programmatic management of your VyOS devices.

- `pyvyos Documentation on Read the Docs
  <https://pyvyos.readthedocs.io/en/latest/>`_ provides detailed instructions
  on the installation, configuration, and operation of the pyvyos library.
- `pyvyos Source Code on GitHub <https://github.com/robertoberto/pyvyos>`_
  allows you to access and contribute to the library's code.
- `pyvyos on PyPI <https://pypi.org/project/pyvyos/>`_ for easy installation
  via pip, the Python package installer. Execute `pip install pyvyos` in your
  terminal to install.


Installation
------------

You can install pyvyos using pip:

.. code-block:: bash

    go install "github.com/ganawaj/go-vyos/vyos"

Getting Started
---------------

Importing and Disabling TLS Verification using Insecure() method of the client
-------------------------------------------------

.. code-block:: none

    import "github.com/ganawaj/go-vyos/vyos"
    client := vyos.NewClient(nil).WithToken("AUTH_KEY").WithURL("https://192.168.0.1").Insecure()

Using API Response Class
------------------------

.. code-block:: none

    type RawResponse struct {
      Success bool        `json:"success,omitempty"`
      Data    interface{} `json:"data,omitempty"`
      Error   string      `json:"error,omitempty"`
    }

Initializing a VyDevice Object
------------------------------

.. code-block:: none

    import (
      "github.com/ganawaj/go-vyos/vyos"
      "os"
    )

    hostname := os.Getenv('VYDEVICE_HOSTNAME')
    port := os.Getenv('VYDEVICE_PORT')
    url := fmt.Sprintf("https://%s:%s", hostname, port)

    apikey := os.Getenv('VYDEVICE_APIKEY')
    verify_ssl := os.Getenv('VYDEVICE_VERIFY_SSL')

    client := vyos.NewClient(nil).WithToken(apikey).WithURL(url)

    if verify_ssl == "false" {
      client = client.Insecure()
    }

Using pyvyos
------------

Configure, then Set
^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: none

    out, resp, err := c.Conf.Set(ctx, "interfaces ethernet eth0 address 192.168.1.1/24")
    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)

Show a Single Object Value
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: none

    out, resp, err := c.Show.Do(ctx, "interfaces dummy dum1 address")
    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)
    fmt.Printf("Data: %v\n", out.Data)

Configure, then Show Object
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: none

    out, resp, err := c.Conf.Get(ctx, "interfaces dummy dum1", nil)
    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)
    fmt.Printf("Data: %v\n", out.Data)

Configure, then Show Multivalue Object
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: none

    options := RetrieveOptions{
        Multivalue: true,
    }

    out, resp, err := c.Conf.Get(ctx, "interfaces dummy dum1", options)
    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)


Configure, then Delete Object
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: none

    out, resp, err := c.Conf.Delete(ctx, "interfaces dummy dum1")
    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)

Configure, then Save
^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: none

    out, resp, err := c.Conf.Save(ctx, "")

    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)

Configure, then Save File
-------------------------

.. code-block:: none

    out, resp, err := c.Conf.Save(ctx, "/config/test300.config")

    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)

Show Object
^^^^^^^^^^^^^^

.. code-block:: none

    out, resp, err := c.Show.Do(ctx, "system image")
    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)
    fmt.Printf("Data: %v\n", out.Data)

Generate Object
^^^^^^^^^^^^^^^^

.. code-block:: none

    out, resp, err := c.Generate.Do(ctx, "pki wireguard key-pair")
    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)
    fmt.Printf("Data: %v\n", out.Data)

Reset Object
^^^^^^^^^^^^^^

.. code-block:: none

    out, resp, err := c.Reset.Do(ctx, "ip bgp 192.0.2.11")
    if err != nil {
        panic("Error: %v", err)
    }

    fmt.Println(out.Success)
    fmt.Printf("Data: %v\n", out.Data)

Configure, then Load File
^^^^^^^^^^^^^^^^^^^^^^^^^^^^

.. code-block:: none

    out, resp, err := c.ConfigFile.Load(ctx, "/config/test300.config")

.. _go-vyos: https://github.com/ganawaj/go-vyos