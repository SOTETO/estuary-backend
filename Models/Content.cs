using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text.Json.Serialization;

namespace estuary_backend.Models
{
    public class Content
    {
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public int Id {  get; set; }

        public string Title { get; set; }

        [JsonIgnore]
        public Workshop Workshop { get; set; }

        //public Guid WorkshopId {  get; set; }

        public List<ContentLink> RelatedContent { get; set; }

        public List<Like> Likes { get; set; }

    }
}
