let PageComponents = [
  {
    type: "Heading",
    disabled: false,
    name: "Heading",
    icon: "text_fields",
    properties: {
      labelSelector: "text"
    }
  },
  {
    type: "Textarea",
    disabled: false,
    name: "Text",
    icon: "wrap_text",
    properties: {}
  },
  {
    type: "HTML",
    disabled: false,
    name: "HTML",
    icon: "code",
    properties: {}
  },
  {
    type: "Quote",
    disabled: false,
    name: "Quote",
    icon: "format_quote",
    properties: {
      labelSelector: "quote"
    }
  },
  {
    type: "Image",
    disabled: false,
    name: "Image",
    icon: "add_a_photo",
    properties: {
      labelSelector: "caption"
    }
  },
  {
    type: "Youtube",
    disabled: false,
    name: "Youtube Video",
    icon: "video_library",
    properties: {
      labelSelector: "videoId"
    }
  },
  {
    type: "Indicator",
    disabled: true,
    name: "Indicator",
    icon: "exposure_plus_2",
    properties: {
      labelSelector: "indicator"
    }
  },
  {
    type: "Transcript",
    disabled: true,
    name: "Transcript",
    icon: "translate",
    properties: {
      labelSelector: "description"
    }
  },
  {
    type: "Footer",
    disabled: false,
    name: "Footer",
    icon: "vertical_align_bottom",
    properties: {
      labelSelector: "value"
    }
  },
  {
    type: "Markdown",
    disabled: false,
    name: "Markdown",
    icon: "code",
    properties: {}
  },
  {
    type: "HTMLEmbed",
    disabled: false,
    name: "HTML Embed",
    icon: "share",
    properties: {
      labelSelector: "embed"
    }
  },
  {
    type: "Legacy",
    disabled: true,
    name: "Legacy Content",
    icon: "lock",
    properties: {}
  },
  {
    type: "FileUpload",
    disabled: false,
    name: "File Upload",
    icon: "cloud_upload",
    properties: {
      eventCategory: "Files",
      labelSelector: "label"
    }
  },
  {
    type: "PDFViewer",
    disabled: false,
    name: "PDF Viewer",
    icon: "picture_as_pdf",
    properties: {
      viewerPath: "",
      eventCategory: "",
      labelSelector: "label"
    }
  }
];

// PDFViewer is Lardwaz block which relies on an external PDF viewer to work.
// It accepts the viewerPath key as path for the viewer e.g => viewerPath: "/viewer/index.html?pdf="

module.exports = PageComponents;

// # NOT IN USE

// {
//   type: "ArticleCard",
//   disabled: true,
//   name: "Article Card 1",
//   icon: "text_fields"
// },
// {
//   type: "Related",
//   disabled: true,
//   name: "Articles",
//   icon: "link",
//   properties: {
//
//   }
// },
// {
//   type: "Gallery",
//   disabled: true,
//   name: "Gallery",
//   icon: "photo_library",
//   properties: {
//
//   }
// }

// # EMAILING BLOCKS

// {
//   type: "MailAd",
//   disabled: true,
//   name: "Mail Ad 1",
//   icon: "text_fields"
// },
// {
//   type: "MailHeader",
//   disabled: true,
//   name: "Mail Header 1",
//   icon: "text_fields"
// },
// {
//   type: "MailJumbotron",
//   disabled: true,
//   name: "Mail Jumbotron 1",
//   icon: "text_fields"
// },
// {
//   type: "MailColumn",
//   disabled: true,
//   name: "Mail Column 1",
//   icon: "text_fields"
// },
// {
//   type: "MailTextHeading",
//   disabled: true,
//   name: "Mail Text Heading 1",
//   icon: "text_fields"
// },
// {
//   type: "MailImageText",
//   disabled: true,
//   name: "Mail Image Text 1",
//   icon: "text_fields"
// },
// {
//   type: "MailArticlesListing",
//   disabled: true,
//   name: "Mail Articles Listing 1",
//   icon: "text_fields"
// }
