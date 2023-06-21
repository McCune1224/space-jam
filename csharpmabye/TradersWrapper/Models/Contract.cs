// {
//     "id": "string",
//   "factionSymbol": "string",
//   "type": "PROCUREMENT",
//   "terms": {
//         "deadline": "2019-08-24T14:15:22Z",
//     "payment": {
//             "onAccepted": 0,
//       "onFulfilled": 0
//     },
//     "deliver": [
//       { }
//     ]
//   },
//   "accepted": false,
//   "fulfilled": false,
//   "expiration": "2019-08-24T14:15:22Z",
//   "deadlineToAccept": "2019-08-24T14:15:22Z"
// }
//

public class Contract
{
    public string? Id { get; set; }
    public string? FactionSymbol { get; set; }
    public string? Type { get; set; }
    public ContractTerms? Terms { get; set; }
    public bool Accepted { get; set; }
    public bool Fulfilled { get; set; }
    public DateTime Expiration { get; set; }
    public DateTime DeadlineToAccept { get; set; }
}



